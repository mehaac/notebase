package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/biozz/wow/notebase/internal/caldav"
	"github.com/biozz/wow/notebase/internal/notebasesync"
	_ "github.com/biozz/wow/notebase/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	root := os.Getenv("NOTES_ROOT")
	if strings.HasPrefix(root, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		root = filepath.Join(dir, root[1:])
	} else if strings.HasPrefix(root, ".") {
		wd, _ := os.Getwd()
		root = filepath.Join(wd, root)
	}
	syncHandler, err := notebasesync.NewHandler(app, root)
	if err != nil {
		app.Logger().Error("error creating sync handler", "error", err)
		return
	}
	caldavHandler := caldav.NewHandler(app, root)

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		syncHandler.Routes(se)
		caldavHandler.Routes(se)

		syncHandler.InitialSync()
		go syncHandler.WatcherManager()

		return se.Next()
	})

	app.OnRecordAfterUpdateSuccess("files").BindFunc(func(e *core.RecordEvent) error {
		syncHandler.OnRecordUpdate(e.Record)
		return e.Next()
	})

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
		Dir:         "./migrations",
	})

	app.RootCmd.AddCommand(syncHandler.SyncCmd())

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
