package gui

import (
	"image"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// NewStreamControl returns the stream widget collection
func NewStreamControl() (*StreamControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if boxErr != nil {
		return nil, boxErr
	}

	optsBox, comboBox, optsErr := newOptionsBox()
	if optsErr != nil {
		return nil, optsErr
	}
	box.PackStart(optsBox, false, false, 0)

	streamWindow, image, streamErr := newStreamWindow()
	if streamErr != nil {
		return nil, streamErr
	}
	box.PackStart(streamWindow, true, true, 0)

	return &StreamControl{rootBox: box, ComboBox: comboBox, image: image}, nil
}

func packDeviceSelector() (*gtk.Label, *gtk.ComboBoxText, error) {
	devicesLabel, labelErr := gtk.LabelNew("Device: ")
	if labelErr != nil {
		return nil, nil, labelErr
	}

	devicesCombo, comboErr := gtk.ComboBoxTextNew()
	if comboErr != nil {
		return nil, nil, comboErr
	}

	devicesCombo.AppendText("DropletGenomics")
	devicesCombo.AppendText("Video file...")

	return devicesLabel, devicesCombo, nil
}

func newOptionsBox() (gtk.IWidget, *gtk.ComboBoxText, error) {
	frame, frameErr := gtk.FrameNew("Device options")
	if frameErr != nil {
		return nil, nil, frameErr
	}

	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if boxErr != nil {
		return nil, nil, boxErr
	}

	devicesLabel, comboBox, comboBoxErr := packDeviceSelector()
	if comboBoxErr != nil {
		return nil, nil, comboBoxErr
	}
	box.PackStart(devicesLabel, false, false, 0)
	box.PackStart(comboBox, false, false, 0)

	frame.Add(box)
	return frame, comboBox, nil
}

func newStreamWindow() (gtk.IWidget, *gtk.Image, error) {
	frame, frameErr := gtk.FrameNew("Device stream")
	if frameErr != nil {
		return nil, nil, frameErr
	}

	image, imgErr := gtk.ImageNew()
	if imgErr != nil {
		return nil, nil, imgErr
	}

	frame.Add(image)
	return frame, image, nil
}

// StreamControl contains the stream window and device controls
type StreamControl struct {
	Control

	// VBox splits the stream window from devices
	rootBox *gtk.Box

	// Device controls
	optionsBox *gtk.Box

	// ComboBox must be accessible
	ComboBox *gtk.ComboBoxText
	image    *gtk.Image

	// Stream frame loader
	pixbufLoader *gdk.PixbufLoader

	device device.Device
}

// Root returns the root widget
func (sw *StreamControl) Root() gtk.IWidget {
	return sw.rootBox
}

// ShowFrame sets an image onto the frame window from an image
func (sw *StreamControl) ShowFrame(frame image.Image) error {
	return showFrame(sw.image, frame)
}
