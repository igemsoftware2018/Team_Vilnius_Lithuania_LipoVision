package device

import "context"

//Device Is a physical device that program is connecting to
type Device interface {

	//Stream Returns a device's video stream that can be cancelled
	Stream(context.Context) <-chan struct{}

	//Available Checks if device is available
	Available() bool
}
