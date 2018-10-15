package gui

import (
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/gotk3/gotk3/gtk"
)

// NewMainControl makes the widget
func NewMainControl(device *device.Device) (*MainControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if boxErr != nil {
		return nil, boxErr
	}

	streamWidget, streamWidgetErr := NewStreamControl(device)
	if streamWidgetErr != nil {
		return nil, streamWidgetErr
	}
	box.PackStart(streamWidget.Root(), true, true, 2)

	pumpAndRegion, regionStream, pumpAndRegionErr := NewPumpAndRegionContainer(device)
	if pumpAndRegionErr != nil {
		return nil, pumpAndRegionErr
	}
	box.PackEnd(pumpAndRegion, false, true, 2)

	box.ShowAll()
	return &MainControl{rootBox: box, StreamControl: streamWidget, RegionStream: regionStream}, nil
}

// MainControl is the root widget of the main window
type MainControl struct {
	Control

	// Contained elements of main window
	StreamControl *StreamControl
	RegionStream  *RegionControl

	// HBox is the hozirontal layout
	// Stream window and device selectors on the left
	// Pump controls on the right
	rootBox *gtk.Box
}

// Root returns the box that contains everything
func (m *MainControl) Root() gtk.IWidget {
	return m.rootBox
}
