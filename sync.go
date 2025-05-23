package main

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gobwas/glob"
	"github.com/goccy/go-yaml"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
	"github.com/syncthing/notify"
)

type SyncHandler struct {
	app       *pocketbase.PocketBase
	root      string
	controlCh chan bool
}

func NewSyncHandler(app *pocketbase.PocketBase, root string) *SyncHandler {
	return &SyncHandler{
		app:  app,
		root: root,
		// It is buffered to send a start signal in the beginning
		controlCh: make(chan bool, 1),
	}
}

func (h *SyncHandler) SyncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Scan notes directory and load markdown files into the files table",
		Run: func(cmd *cobra.Command, args []string) {
			go h.JobManager()
		},
	}
}

func (h *SyncHandler) Routes(se *core.ServeEvent) {
	syncGroup := se.Router.Group("/sync")
	syncGroup.Bind(apis.RequireSuperuserAuth())
	syncGroup.GET("/start", func(e *core.RequestEvent) error {
		h.controlCh <- true
		return nil
	})
	syncGroup.GET("/stop", func(e *core.RequestEvent) error {
		h.controlCh <- false
		return nil
	})
}

type NotebaseConfig struct {
	ClearOnStartup bool     `yaml:"clear_on_startup"`
	Exclude        []string `yaml:"exclude"`
}

// syncJob loads markdown files into an SQLite database and synchronizes
//
//	the changes between the files system and the database.
//
// There are several parts to the sync job:
//   - it clears the database on startup and heavily rely on this during development
//   - it has a bunch of other settings, like notes root dir and exclude patterns
//   - it starts a filesystem watcher early, because I am going to go through all
//     the directories upon initial scan anyway
//   - all the parsing is done in parallel in goroutines
//   - after parsing the results are sent to the saving process
//     which also works in parallel in goroutine in batches
func (h *SyncHandler) job(stopCh <-chan struct{}, doneCh chan<- struct{}) {
	data, err := os.ReadFile(path.Join(h.root, ".notebase.yml"))
	if err != nil {
		h.app.Logger().Error("error reading config", "error", err)
		return
	}
	config := NotebaseConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		h.app.Logger().Error("error parsing config", "error", err)
		return
	}

	startTime := time.Now()

	if config.ClearOnStartup {
		// TODO: add metrics to find out if it works
		// app.DB().NewQuery("PRAGMA synchronous = OFF").Execute()
		// app.DB().NewQuery("PRAGMA journal_mode = MEMORY").Execute()

		_, err := h.app.DB().NewQuery("DELETE FROM files").Execute()
		if err != nil {
			h.app.Logger().Error("unable to clear files table", err)
			return
		}
	}

	patterns := make([]glob.Glob, 0, len(config.Exclude))
	for _, pattern := range config.Exclude {
		g, err := glob.Compile(pattern)
		if err != nil {
			h.app.Logger().Error("Invalid glob pattern", "pattern", pattern, "error", err)
			continue
		}
		patterns = append(patterns, g)
	}

	fileChanges := make(chan notify.EventInfo, 1)
	watchPath := path.Join(h.root, "/...")
	if err := notify.Watch(watchPath, fileChanges, notify.All); err != nil {
		h.app.Logger().Error("error watching files", "error", err)
	}
	defer notify.Stop(fileChanges)
	go h.fileWatcher(fileChanges, patterns)

	filesChan := make(chan string, 200)
	resultsChan := make(chan File, 200)

	var parseWg sync.WaitGroup
	var saveWg sync.WaitGroup

	// TODO: move to config
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		parseWg.Add(1)
		go func() {
			defer parseWg.Done()
			for wp := range filesChan {
				data, err := parse(h.root, wp)
				if err != nil {
					h.app.Logger().Error("error parsing file", "path", wp, "error", err)
					continue
				}
				resultsChan <- data
			}
		}()
	}

	// TODO: move to config
	const batchSize = 200
	saveWg.Add(1)
	go func() {
		defer saveWg.Done()
		batch := make([]File, 0, batchSize)

		processBatch := func() {
			if len(batch) == 0 {
				return
			}

			filesCol, _ := h.app.FindCollectionByNameOrId("files")
			err := h.app.RunInTransaction(func(txApp core.App) error {
				for _, data := range batch {
					fileRec := core.NewRecord(filesCol)
					fillFileRecFromData(fileRec, data)
					// SaveNoValidate probably speeds things up a bit
					if err := txApp.SaveNoValidate(fileRec); err != nil {
						return err
					}
				}
				return nil
			})

			if err != nil {
				h.app.Logger().Error("error saving batch", "error", err)
				return
			}
			// Clear batch but reuse the underlying array
			batch = batch[:0]
		}

		for data := range resultsChan {
			batch = append(batch, data)
			if len(batch) >= batchSize {
				processBatch()
			}
		}
		processBatch()
	}()

	err = filepath.WalkDir(h.root, func(walkPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		p, _ := filepath.Rel(h.root, walkPath)

		if isExcluded(patterns, p) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if filepath.Ext(walkPath) == ".md" {
			filesChan <- walkPath
		}
		return nil
	})

	if err != nil {
		h.app.Logger().Error("unable to walk files", err)
	}

	close(filesChan)
	parseWg.Wait()
	close(resultsChan)
	saveWg.Wait()

	elapsedTime := time.Since(startTime)
	h.app.Logger().Info("Initial sync complete", "elapsed", elapsedTime.String())

	// Endless loop required to keep running and restart the job
	for {
		select {
		case <-stopCh:
			doneCh <- struct{}{}
			return
		}
	}
}

func (h *SyncHandler) JobManager() {
	var internalStopCh chan struct{}
	var doneCh chan struct{}

	h.controlCh <- true

	for {
		select {
		case cmd, ok := <-h.controlCh:
			if !ok {
				// Channel closed, stop sync job if running and return
				if internalStopCh != nil {
					close(internalStopCh)
					<-doneCh
				}
				return
			}

			if cmd {
				if internalStopCh == nil {
					internalStopCh = make(chan struct{})
					doneCh = make(chan struct{})
					go h.job(internalStopCh, doneCh)
				}
			} else {
				if internalStopCh != nil {
					close(internalStopCh)
					<-doneCh
					internalStopCh = nil
					doneCh = nil
				}
			}
		}
	}
}

func (h *SyncHandler) fileWatcher(watcher chan notify.EventInfo, patterns []glob.Glob) {
	for {
		select {
		case ei := <-watcher:
			relPath, _ := filepath.Rel(h.root, ei.Path())
			if isExcluded(patterns, relPath) {
				continue
			}
			h.app.Logger().Info("file watcher event", "event", ei.Event().String(), "path", ei.Path())
			if filepath.Ext(ei.Path()) != ".md" {
				continue
			}
			if ei.Event() == notify.Create {
				fstat, _ := os.Stat(ei.Path())
				if fstat.IsDir() {
					continue
				}
			}
			h.handleFSNotifyEvent(ei)
		}
	}
}

type File struct {
	FrontMatter     yaml.MapSlice
	FrontMatterJSON string
	Content         string
	AbsPath         string
	RelPath         string
	Slug            string
	Links           []string
	Tags            []string
}

var (
	linkRegex = regexp.MustCompile(`\[\[(\S*)\]\]`)
)

func parse(rootPath string, curPath string) (File, error) {
	relPath, _ := filepath.Rel(rootPath, curPath)
	fileName := filepath.Base(relPath)
	slug := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	content, err := os.ReadFile(curPath)
	if err != nil {
		return File{}, err
	}

	contentStr := string(content)

	data := File{
		AbsPath:     curPath,
		RelPath:     relPath,
		Slug:        slug,
		FrontMatter: yaml.MapSlice{},
		Links:       []string{},
		Tags:        []string{},
	}

	// Extract frontmatter
	frontMatter, mainContent := extractFrontMatter(contentStr)

	if len(frontMatter) > 0 {
		yaml.Unmarshal([]byte(frontMatter), &data.FrontMatter)
		jsonBytes, err := yaml.MarshalWithOptions(data.FrontMatter, yaml.JSON())
		if err != nil {
			data.FrontMatterJSON = "{}"
		} else {
			data.FrontMatterJSON = string(jsonBytes)
		}
	}

	// Store content
	data.Content = mainContent

	// Find all links in one pass
	linkMap := make(map[string]bool)
	linkMatches := linkRegex.FindAllStringSubmatch(contentStr, -1)
	for _, match := range linkMatches {
		if len(match[1]) > 0 {
			linkMap[match[1]] = true
		}
	}

	// Convert link map to slice
	for link := range linkMap {
		data.Links = append(data.Links, link)
	}

	return data, nil
}

type WatcherEvent struct {
	EventType string
	Path      string
}

func (h *SyncHandler) handleFSNotifyEvent(event notify.EventInfo) {
	switch event.Event() {
	case notify.Create, notify.Write:
		data, err := parse(h.root, event.Path())
		if err != nil {
			h.app.Logger().Error("unable to parse", err)
			return
		}

		if event.Event() == notify.Create {
			if data.Slug == "Untitled" {
				// thi is probably a placeholder file, created by Obsidian, ignore it
				return
			}
			if err := createFile(h.app, data); err != nil {
				h.app.Logger().Error("unable to create file", err)
			}
			return
		}

		if event.Event() == notify.Write {
			if err := h.updateFile(data); err != nil {
				h.app.Logger().Error("unable to write file", err)
			}
			return
		}
	case notify.Rename, notify.Remove:
		if err := h.deleteFile(event.Path()); err != nil {
			h.app.Logger().Error("unable to delete file", err)
		}
	}
}

func createFile(app *pocketbase.PocketBase, data File) error {
	filesCol, err := app.FindCollectionByNameOrId("files")
	if err != nil {
		return err
	}
	fileRec := core.NewRecord(filesCol)
	fillFileRecFromData(fileRec, data)
	if err := app.Save(fileRec); err != nil {
		app.Logger().Error("Error saving file record", "error", err)
		return err
	}
	return nil
}

func fillFileRecFromData(fileRec *core.Record, data File) {
	fileRec.Load(map[string]any{
		"path":        data.RelPath,
		"slug":        data.Slug,
		"content":     data.Content,
		"frontmatter": data.FrontMatterJSON,
	})
}

func (h *SyncHandler) updateFile(data File) error {
	fileRec, err := h.app.FindFirstRecordByData("files", "path", data.RelPath)
	if err != nil {
		return err
	}
	fillFileRecFromData(fileRec, data)
	if err := h.app.Save(fileRec); err != nil {
		h.app.Logger().Error("Error updating file record", "error", err)
		return err
	}
	return nil
}

func (h *SyncHandler) deleteFile(path string) error {
	fileRec, err := h.app.FindFirstRecordByData("files", "path", path)
	if err != nil {
		return err
	}
	if err := h.app.Delete(fileRec); err != nil {
		h.app.Logger().Error("Error deleting file record", "error", err)
		return err
	}
	return nil
}
