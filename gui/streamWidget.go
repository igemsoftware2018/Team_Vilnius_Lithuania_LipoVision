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

	optsBox, optsErr := newOptionsBox()
	if optsErr != nil {
		return nil, optsErr
	}
	box.PackStart(optsBox, false, false, 0)

	streamWindow, streamErr := newStreamWindow()
	if streamErr != nil {
		return nil, streamErr
	}
	box.PackStart(streamWindow, false, true, 0)

	return &StreamWidget{rootBox: box}, nil
}

func newOptionsBox() (gtk.IWidget, error) {
	frame, frameErr := gtk.FrameNew("Device options")
	if frameErr != nil {
		return nil, frameErr
	}

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
	box.PackStart(devicesCombo, false, false, 0)

	frame.Add(box)
	return frame, nil
}

func newStreamWindow() (gtk.IWidget, error) {
	frame, frameErr := gtk.FrameNew("Device stream")
	if frameErr != nil {
		return nil, frameErr
	}

	image, imgErr := gtk.ImageNewFromFile("template-intersection.png")
	if imgErr != nil {
		return nil, imgErr
	}

	frame.Add(image)
	return frame, nil
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
