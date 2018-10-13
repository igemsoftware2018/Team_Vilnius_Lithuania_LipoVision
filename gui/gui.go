package gui

import "log"
import "github.com/mattn/go-gtk/gtk"

func init() {
	gtk.Init(nil)
}

func Compose() {
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	if win == nil {
		log.Fatal("Unable to create window")
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l := gtk.NewLabel("Hello, gotk3!")
	if l == nil {
		log.Fatal("Unable to create label")
	}

	win.Add(l)
	win.SetDefaultSize(800, 600)
	win.ShowAll()

}
