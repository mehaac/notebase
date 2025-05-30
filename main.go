package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/biozz/wow/notebase/internal/caldav"
	"github.com/biozz/wow/notebase/internal/config"
	"github.com/biozz/wow/notebase/internal/notebasesync"
	_ "github.com/biozz/wow/notebase/migrations"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	root := os.Getenv("NOTES_ROOT")
	superuserEmail := os.Getenv("SUPERUSER_EMAIL")
	superuserPassword := os.Getenv("SUPERUSER_PASSWORD")

	if strings.HasPrefix(root, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		root = filepath.Join(dir, root[1:])
	} else if strings.HasPrefix(root, ".") {
		wd, _ := os.Getwd()
		root = filepath.Join(wd, root)
	}

	conf, err := config.Load(root)
	if err != nil {
		app.Logger().Error("error loading config", "error", err)
		return
	}

	syncHandler, err := notebasesync.NewHandler(app, root, &conf)
	if err != nil {
		app.Logger().Error("error creating sync handler", "error", err)
		return
	}
	caldavHandler := caldav.NewHandler(app, root, &conf)

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.InstallerFunc = CustomInstallerFunc(superuserEmail, superuserPassword)

		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		syncHandler.Routes(se)
		caldavHandler.Routes(se)

		// TODO: run this in a goroutine, but make sure that watcher is not running, while syncing
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

// This is basically a copy-paste of the `superuser` command from pocketbase
func CustomInstallerFunc(superuserEmail, superuserPassword string) func(app core.App, systemSuperuser *core.Record, baseURL string) error {
	return func(app core.App, systemSuperuser *core.Record, baseURL string) error {
		if superuserEmail == "" || superuserPassword == "" {
			return fmt.Errorf("SUPERUSER_USERNAME or SUPERUSER_PASSWORD are empty, skipping superuser creation")
		}
		if is.EmailFormat.Validate(superuserEmail) != nil {
			return errors.New("missing or invalid email address")
		}

		superusersCol, err := app.FindCachedCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return fmt.Errorf("failed to fetch %q collection: %w", core.CollectionNameSuperusers, err)
		}

		superuser, err := app.FindAuthRecordByEmail(superusersCol, superuserEmail)
		if err != nil {
			superuser = core.NewRecord(superusersCol)
		}

		superuser.SetEmail(superuserEmail)
		superuser.SetPassword(superuserPassword)

		if err := app.Save(superuser); err != nil {
			return fmt.Errorf("Failed to upsert superuser account: %w", err)
		}

		return nil
	}
}
