package notebasesync

import (
	"io/fs"
	"path/filepath"
	"sync"
	"time"

	"github.com/biozz/wow/notebase/internal/utils"
	"github.com/pocketbase/pocketbase/core"
)

func (h *SyncHandler) InitialSync() {
	startTime := time.Now()

	if h.config.ClearOnStartup {
		_, err := h.app.DB().NewQuery("DELETE FROM files").Execute()
		if err != nil {
			h.app.Logger().Error("unable to clear files table", err)
			return
		}
	}

	filesChan := make(chan string, h.config.SyncBatchSize)
	resultsChan := make(chan File, h.config.SyncBatchSize)

	var parseWg sync.WaitGroup
	var saveWg sync.WaitGroup

	// Usually 5 workers is enough to parse about 4k files in <300ms.
	// More workers will bottleneck saver process.
	for i := 0; i < h.config.SyncWorkers; i++ {
		parseWg.Add(1)
		go h.parserManager(&parseWg, filesChan, resultsChan)
	}

	// We don't need workers here, because SQLite runs in a single thread
	saveWg.Add(1)
	go h.saverManager(&saveWg, resultsChan)

	err := filepath.WalkDir(h.root, func(walkPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		p, _ := filepath.Rel(h.root, walkPath)

		if utils.IsExcluded(h.excludePatters, p) {
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
}

func (h *SyncHandler) parserManager(parseWg *sync.WaitGroup, filesChan <-chan string, resultsChan chan<- File) {
	defer parseWg.Done()
	for wp := range filesChan {
		data, err := parse(h.root, wp)
		if err != nil {
			h.app.Logger().Error("error parsing file", "path", wp, "error", err)
			continue
		}
		resultsChan <- data
	}
}

func (h *SyncHandler) saverManager(saveWg *sync.WaitGroup, resultsChan <-chan File) {
	defer saveWg.Done()
	batch := make([]File, 0, h.config.SyncBatchSize)

	for data := range resultsChan {
		batch = append(batch, data)
		if len(batch) >= h.config.SyncBatchSize {
			h.saver(batch)
		}
	}
	h.saver(batch)
}

func (h *SyncHandler) saver(batch []File) {
	if len(batch) == 0 {
		return
	}

	filesCol, _ := h.app.FindCollectionByNameOrId("files")
	err := h.app.RunInTransaction(func(txApp core.App) error {
		for _, data := range batch {
			fileRec := core.NewRecord(filesCol)
			data.Origin = "init"
			data.Version = utils.GetVersion()
			fillFileRecFromData(fileRec, data)
			// SaveNoValidate probably speeds things up a bit
			if err := txApp.SaveNoValidate(fileRec); err != nil {
				return err
			}
			utils.SetFileXAttrs(data.AbsPath, utils.XAttrs{Version: data.Version, Origin: data.Origin})
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
