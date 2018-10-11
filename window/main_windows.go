package window

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// NewMain creates the main window
func NewMain() (*Main, error) {
	walk.Resources.SetRootDirPath(".")

	type Mode struct {
		Name  string
		Value ImageViewMode
	}

	modes := []Mode{
		{"ImageViewModeIdeal", ImageViewModeIdeal},
		{"ImageViewModeCorner", ImageViewModeCorner},
		{"ImageViewModeCenter", ImageViewModeCenter},
		{"ImageViewModeShrink", ImageViewModeShrink},
		{"ImageViewModeZoom", ImageViewModeZoom},
		{"ImageViewModeStretch", ImageViewModeStretch},
	}

	var widgets []Widget

	for _, mode := range modes {
		widgets = append(widgets,
			Label{
				Text: mode.Name,
			},
			ImageView{
				Background: SolidColorBrush{Color: walk.RGB(255, 191, 0)},
				Image:      "template-intersection.png",
				Margin:     10,
				Mode:       mode.Value,
			},
		)
	}

	win := MainWindow{
		Title:    "Walk ImageView Example",
		Size:     Size{400, 600},
		Layout:   Grid{Columns: 2},
		Children: widgets,
	}

	return Main{
		window: win,
	}, nil
}

// Main defines the main window
type Main struct {
	Window

	events map[string]chan Event
	window *MainWindow
}

// Run starts the main window and event loop
func (w Main) Run() {
	w.window.Run()
}

// Subscribe provides a new channel of events
func (w Main) Subscribe(eventName string) <-chan Event {
	newEvent := make(chan Event, EventQueueLength)
	w.events[eventName] = newEvent
	return newEvent
}
