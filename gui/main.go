package gui

import "github.com/gotk3/gotk3/gtk"

// NewMainControl makes the widget
func NewMainControl() (*MainControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if boxErr != nil {
		return nil, boxErr
	}

	streamWidget, streamWidgetErr := NewStreamControl()
	if streamWidgetErr != nil {
		return nil, streamWidgetErr
	}
	box.PackStart(streamWidget.Root(), true, true, 2)

	pumpAndRegion, pumpAndRegionErr := NewPumpAndRegionContainer()
	if pumpAndRegionErr != nil {
		return nil, pumpAndRegionErr
	}
	box.PackEnd(pumpAndRegion, false, true, 2)

	box.ShowAll()
	return &MainControl{rootBox: box}, nil
}

// MainControl is the root widget of the main window
type MainControl struct {
	Control

	// HBox is the hozirontal layout
	// Stream window and device selectors on the left
	// Pump controls on the right
	rootBox *gtk.Box
}

// Root returns the box that contains everything
func (m *MainControl) Root() gtk.IWidget {
	return m.rootBox
}
