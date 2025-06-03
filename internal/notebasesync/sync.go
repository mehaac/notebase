package notebasesync

import (
	"fmt"
	"path"

	"github.com/biozz/wow/notebase/internal/config"
	"github.com/gobwas/glob"
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
	excludePatters []glob.Glob
	conf           *config.NotebaseConfig

	// watcher
	fileChanges chan notify.EventInfo
}

func NewHandler(app *pocketbase.PocketBase, root string, conf *config.NotebaseConfig) (*SyncHandler, error) {

	patterns := make([]glob.Glob, 0, len(conf.Exclude))
	for _, pattern := range conf.Exclude {
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
		conf: conf,
		// It is buffered to send a start signal in the beginning
		controlCh:      make(chan bool, 1),
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
	syncGroup.GET("/restart", func(e *core.RequestEvent) error {
		h.controlCh <- false
		h.InitialSync()
		h.controlCh <- true
		return nil
	})
}
