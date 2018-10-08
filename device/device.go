//go:generate mockgen -destination mock_device/mock_device.go github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device Device,Frame,Client

package device

import (
	"context"
	"image"
)

// Frame Defines a frame structure that is a part of the stream
type Frame interface {

	// Frame Gets the underlying frame
	Frame() image.Image

	// SkippedFrame provides context.Done() like method of
	// communicating that a frame should be skipped
	Skip() <-chan struct{}
}

// Device Is the physical or virtual device that this program is connecting to,
// It can have whatever parts it wants
type Device interface {

	//S tream Returns a device's video stream that can be cancelled
	Stream(context.Context) <-chan Frame

	// Available Checks if device is available
	Available() bool
}
