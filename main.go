package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/biozz/wow/notebase/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		initCaldavRoutes(app, se)

		root, _ := app.RootCmd.Flags().GetString("root")
		go syncJob(app, root)

		return se.Next()
	})

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
		Dir:         "./migrations",
	})

	app.RootCmd.PersistentFlags().String("root", "", "a path to notes root, which has .notebase.yml")
	app.RootCmd.AddCommand(syncCmd(app))

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
