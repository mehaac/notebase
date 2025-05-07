package main

import "github.com/pocketbase/pocketbase"

func getSetting(app *pocketbase.PocketBase, key string) string {
	rec, err := app.FindFirstRecordByData("settings", "key", key)
	if err != nil {
		app.Logger().Error("unable to find setting", err)
		return ""
	}
	return rec.GetString("value")
}
