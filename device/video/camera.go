package video

import (
	"errors"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

// Camera defines a "camera" for the video device
type Camera struct {
	device.Client

	FrameRate int
}

// Invoke performs a configuration change on the camera
func (camera *Camera) Invoke(invoke device.ClientInvocation, data interface{}) error {
	if invoke != device.CameraSetFrameRate {
		return errors.New("this camera operation is unsupported")
	}
	camera.FrameRate = data.(int)
	return nil
}
