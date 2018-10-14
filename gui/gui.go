package gui

import "github.com/gotk3/gotk3/gtk"

func init() {
	gtk.Init(nil)
}

// Widget interface for this project
type Widget interface {
	// Root returns gotk3 type of widget
	// This widget is root for all represented components of this widget
	Root() gtk.IWidget
}
