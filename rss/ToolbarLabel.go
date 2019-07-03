package rss

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	//"fyne.io/fyne/canvas" // XXX
)

type ToolbarLabel struct {
	label *widget.Label
}

// Implementing widget.ToolbarItem interface
func (tl *ToolbarLabel) ToolbarObject() fyne.CanvasObject {
	return tl.label
}

func (tl *ToolbarLabel) SetText(text string) {
  tl.label.SetText(text)
}

func NewToolbarLabel(text string) *ToolbarLabel {
	tl := &ToolbarLabel{
		label: widget.NewLabel(text),
	}
	return tl
}

// TODO: Implement nice destroy methods ???