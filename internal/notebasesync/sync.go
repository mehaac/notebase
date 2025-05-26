package notebasesync

import (
	"fmt"
	"os"
	"path"

	"github.com/gobwas/glob"
	"github.com/goccy/go-yaml"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
	"github.com/syncthing/notify"
)

type SyncHandler struct {
	app            *pocketbase.PocketBase
	root           string
	controlCh      chan bool
	config         NotebaseConfig
	excludePatters []glob.Glob

	// watcher
	fileChanges chan notify.EventInfo
}

type NotebaseConfig struct {
	ClearOnStartup bool     `yaml:"clear_on_startup"`
	Exclude        []string `yaml:"exclude"`
	SyncWorkers    int      `yaml:"sync_workers"`
	SyncBatchSize  int      `yaml:"sync_batch_size"`
}

func NewHandler(app *pocketbase.PocketBase, root string) (*SyncHandler, error) {
	data, err := os.ReadFile(path.Join(root, ".notebase.yml"))
	if err != nil {
		return nil, err
	}
	config := NotebaseConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config", err)
	}

	if config.SyncBatchSize == 0 {
		config.SyncBatchSize = 200
	}
	if config.SyncWorkers == 0 {
		config.SyncWorkers = 5
	}

	patterns := make([]glob.Glob, 0, len(config.Exclude))
	for _, pattern := range config.Exclude {
		g, err := glob.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("Invalid glob pattern", err)
		}
		patterns = append(patterns, g)
	}

	fileChanges := make(chan notify.EventInfo, 1)
	watchPath := path.Join(root, "/...")
	if err := notify.Watch(watchPath, fileChanges, notify.All); err != nil {
		return nil, fmt.Errorf("error starting file watcher", err)
	}

	return &SyncHandler{
		app:  app,
		root: root,
		// It is buffered to send a start signal in the beginning
		controlCh:      make(chan bool, 1),
		config:         config,
		excludePatters: patterns,
		fileChanges:    fileChanges,
	}, nil
}

func (h *SyncHandler) SyncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Scan notes directory and load markdown files into the files table",
		Run: func(cmd *cobra.Command, args []string) {
			h.InitialSync()
			go h.WatcherManager()
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
