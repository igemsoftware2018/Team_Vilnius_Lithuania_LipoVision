package video

import (
	"errors"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

// Camera defines a "camera" for the video device
type Camera struct {
	device.Client

	FrameRate float64
}

// Invoke performs a configuration change on the camera
func (camera *Camera) Invoke(invoke device.ClientInvocation, data float64) error {
	if invoke != device.CameraSetFrameRate {
		return errors.New("this camera operation is unsupported")
	}
	camera.FrameRate = data
	return nil
}
