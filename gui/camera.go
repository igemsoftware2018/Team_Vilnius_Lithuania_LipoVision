package gui

import "github.com/gotk3/gotk3/gtk"

// NewCameraConrol returns the camera controls collection
func NewCameraConrol() (*CameraControl, error) {
	frame, frameErr := gtk.FrameNew("Camera contols")
	if frameErr != nil {
		return nil, frameErr
	}

	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if boxErr != nil {
		return nil, boxErr
	}

	illuminationScale, exposureScale, pckErr := packNewScales(box)
	if pckErr != nil {
		return nil, pckErr
	}

	autoAdjButton, audoAdjButtonErr := gtk.ButtonNewWithLabel("Auto Adjust")
	if audoAdjButtonErr != nil {
		return nil, audoAdjButtonErr
	}
	box.PackStart(autoAdjButton, true, false, 5)

	frame.Add(box)
	return &CameraControl{rootFrame: frame,
		AutoAdjButton:     autoAdjButton,
		IlluminationScale: illuminationScale,
		ExposureScale:     exposureScale}, nil
}

func packNewScales(box *gtk.Box) (*gtk.Scale, *gtk.Scale, error) {
	illLabel, illLabelErr := gtk.LabelNew("Illumination")
	if illLabelErr != nil {
		return nil, nil, illLabelErr
	}
	box.PackStart(illLabel, false, false, 10)

	illuminationScale, illuminationScaleErr := gtk.ScaleNewWithRange(
		gtk.ORIENTATION_HORIZONTAL,
		0, 100, 1)
	if illuminationScaleErr != nil {
		return nil, nil, illuminationScaleErr
	}
	box.PackStart(illuminationScale, true, false, 0)

	expLabel, expLabelErr := gtk.LabelNew("Exposure")
	if expLabelErr != nil {
		return nil, nil, expLabelErr
	}
	box.PackStart(expLabel, false, false, 10)

	exposureScale, exposureScaleErr := gtk.ScaleNewWithRange(
		gtk.ORIENTATION_HORIZONTAL,
		0, 100, 1)
	if exposureScaleErr != nil {
		return nil, nil, exposureScaleErr
	}
	box.PackStart(exposureScale, true, false, 0)

	return illuminationScale, exposureScale, nil
}

// CameraControl is a collection of controls for the camera
type CameraControl struct {
	Control

	AutoAdjButton     *gtk.Button
	IlluminationScale *gtk.Scale
	ExposureScale     *gtk.Scale

	rootFrame *gtk.Frame
}

// Root returns the root element of these controls
func (cc *CameraControl) Root() gtk.IWidget {
	return cc.rootFrame
}
