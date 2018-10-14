package gui

import "github.com/gotk3/gotk3/gtk"

// NewMainWidget makes the widget
func NewMainWidget() (*MainWidget, error) {

	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if boxErr != nil {
		return nil, boxErr
	}

	streamWidget, streamWidgetErr := NewStreamWidget()
	if streamWidgetErr != nil {
		return nil, streamWidgetErr
	}
	box.PackStart(streamWidget.Root(), true, true, 0)

	pumpControlsWidget, pumpControlsWidgetErr := NewPumpControlsWidget()
	if pumpControlsWidgetErr != nil {
		return nil, pumpControlsWidgetErr
	}
	box.PackEnd(pumpControlsWidget.Root(), false, true, 0)

	box.ShowAll()
	return &MainWidget{rootBox: box}, nil
}

// MainWidget is the root widget of the main window
type MainWidget struct {
	Widget

	// HBox is the hozirontal layout
	// Stream window and device selectors on the left
	// Pump controls on the right
	rootBox *gtk.Box
}

// Widget returns the box that contains everything
func (m *MainWidget) Root() gtk.IWidget {
	return m.rootBox
}
