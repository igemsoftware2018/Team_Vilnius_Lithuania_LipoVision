package gui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// NewStreamWidget returns the stream widget collection
func NewStreamWidget() (*StreamWidget, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if boxErr != nil {
		return nil, boxErr
	}

	return &StreamWidget{rootBox: box}, nil
}

func newOptionsBox() (gtk.IWidget, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if boxErr != nil {
		return nil, boxErr
	}

	devicesLabel, labelErr := gtk.LabelNew("Devices: ")
	if labelErr != nil {
		return nil, labelErr
	}
	box.PackStart(devicesLabel, false, false, 0)

	devicesCombo, comboErr := gtk.ComboBoxNew()
	if comboErr != nil {
		return nil, comboErr
	}
	devicesCombo.SetModel(nil)
	box.PackStart(devicesCombo, false, false, 0)

	return box, nil
}

// StreamWidget contains the stream window and device controls
type StreamWidget struct {
	Widget

	// VBox splits the stream window from devices
	rootBox *gtk.Box

	// Device controls
	optionsBox *gtk.Box

	// Stream
	image *gtk.Image

	// Stream frame loader
	pixbufLoader *gdk.PixbufLoader
}

func (sw *StreamWidget) Root() gtk.IWidget {
	return sw.rootBox
}
