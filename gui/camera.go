package gui

import "github.com/gotk3/gotk3/gtk"

func NewCameraConrol() (*CameraControl, error) {
	frame, frameErr := gtk.FrameNew("Camera contols")
	if frameErr != nil {
		return nil, frameErr
	}

	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if boxErr != nil {
		return nil, boxErr
	}

	illuminationScale, illuminationScaleErr := gtk.ScaleNewWithRange(
		gtk.ORIENTATION_HORIZONTAL,
		0, 100, 1)
	if illuminationScaleErr != nil {
		return nil, illuminationScaleErr
	}
	box.PackStart(illuminationScale, true, false, 5)

	exposureScale, exposureScaleErr := gtk.ScaleNewWithRange(
		gtk.ORIENTATION_HORIZONTAL,
		0, 100, 1)
	if exposureScaleErr != nil {
		return nil, exposureScaleErr
	}
	box.PackStart(exposureScale, true, false, 5)

	autoAdjButton, audoAdjButtonErr := gtk.ButtonNewWithLabel("Auto Adjust")
	if audoAdjButtonErr != nil {
		return nil, audoAdjButtonErr
	}
	box.PackStart(autoAdjButton, true, false, 5)

	frame.Add(box)
	return &CameraControl{rootFrame: frame}, nil
}

type CameraControl struct {
	Control

	rootFrame *gtk.Frame
}

func (cc *CameraControl) Root() gtk.IWidget {
	return cc.rootFrame
}
