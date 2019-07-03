// Package main launches the RSS reader example directly
package main

import (
	"fyne.io/fyne/app"

	"github.com/fyne-io/examples/img/icon"
	"github.com/fyne-io/examples/rss"
)

func main() {
	app := app.New()
	app.SetIcon(icon.TextEditorBitmap) // FIXME - icon

	rss.Show(app)
	app.Run()
}
