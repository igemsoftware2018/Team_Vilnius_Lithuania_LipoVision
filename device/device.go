package device

type Camera interface {
	StreamAvailable() bool

	//
	ProcessFrames(bool) error
}

type PumpController interface {
}

//Device Is a physical device that program is connecting to
type Device interface {

	//Available Checks if device is available
	Available() bool

	//Camera Gets Device's camera
	Camera() Camera

	//PumpController Gets Device's pump contoller
	PumpController() PumpController
}
