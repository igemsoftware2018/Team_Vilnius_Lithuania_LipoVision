// +build !windows

package window

import (
	"github.com/gotk3/gotk3/gtk"
)

// init gtk
func init() {
	gtk.Init(nil)
}

// NewMain creates the main application window
func NewMain() (*Main, error) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	win.SetTitle("LipoVision")
	win.SetDefaultSize(800, 600)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		return nil, err
	}
	win.Add(l)

	return &Main{
		window: win,
		label:  l,
		events: make(map[string]Event),
	}, nil
}

// Main is the main application window
type Main struct {
	Window

	events map[string]chan Event
	window *gtk.Window
	label  *gtk.Label
}

// Run starts main loop, blocks
func (w Main) Run() {
	w.window.ShowAll()
	gtk.Main()
}

// Subscribe provides a new channel of events
func (w *Main) Subscribe(eventName string) <-chan Event {
	newEvent := make(chan Event, EventQueueLength)
	w.events[eventName] = newEvent
	return newEvent
}
