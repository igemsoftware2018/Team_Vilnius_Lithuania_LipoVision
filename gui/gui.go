package gui

import "log"
import "github.com/gotk3/gotk3/gtk"

func init() {
	gtk.Init(nil)
}

func Compose() {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window")
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label")
	}

	win.Add(l)
	win.SetDefaultSize(800, 600)
	win.ShowAll()

	gtk.Main()
}
