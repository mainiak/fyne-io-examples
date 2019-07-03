package rss

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/examples/img/icon"
)

func setLocationHome(tl *ToolbarLabel) {
	tl.SetText("Feed List")
}

func setLocationNewFeed(tl *ToolbarLabel) {
	tl.SetText("Add new item")
}

func setLocationFeedList(tl *ToolbarLabel, feedName string) {
	tl.SetText(feedName)
}

// Show loads a new RSS reader
func Show(app fyne.App) {
	window := app.NewWindow("RSS Reader")
	window.SetIcon(icon.TextEditorBitmap) // FIXME - icon

	location := NewToolbarLabel(" <-- Please add new RSS/Atom feed by clicking on (+) icon")
	// setLocationHome(location) // can't be here immediately after creation of ToolbarLabel

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			setLocationHome(location)
        }),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			setLocationNewFeed(location)
		}),
		widget.NewToolbarSeparator(),
		location,
	)


	status := widget.NewHBox( //layout.NewSpacer(),
		widget.NewLabel("Bottom bar"))

	content := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, status, nil, nil),
		toolbar,
		status,
		widget.NewScrollContainer(widget.NewButtonWithIcon("", theme.WarningIcon(), func() {}) ),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(480, 320))

	window.Show()
}
