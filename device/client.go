package device

// ClientInvocation is the Client.Invoke() instruction.
// You can define new ones and integrate them easily.
type ClientInvocation int

const (

	// PumpSetTargetVolume sets the 'target' volume that has to be pushed through
	// Input is a float64 in uL.
	PumpSetTargetVolume ClientInvocation = iota

	//PumpReset resets indicated pump to starting position. Input can be nil
	PumpReset

	// PumpToggleWithdrawInfuse allows reversing of direction of the pump,
	// enabling you to pull liquids through the system instead of pushing them through.
	// Input is bool, true is push, false is pull.
	PumpToggleWithdrawInfuse

	// PumpSetVolume sets the volume per second target.
	// Input is a float64 in uL/s.
	PumpSetVolume

	// PumpToggle turns the pump on or off. Input is bool.
	PumpToggle

	// PumpRefresh asks the device for it's parameters and updates values.
	// Input can be nil.
	PumpRefresh

	// PumpPurge instructs a pump to push/pull as fast is possible, used for flushing.
	PumpPurge

	// CameraSetIllumination instructs the camera on how much illumination to apply.
	// Input is and int 0 - 100 in percent.
	CameraSetIllumination

	// CameraSetExposure instructs the camera on how much exposure to apply.
	// Input is and int 0 - 100 in percent.
	CameraSetExposure

	// CameraSetFrameRate instructs the camera on what framerate to send frames.
	// Input is int in fps.
	CameraSetFrameRate

	// CameraAutoAdjust instructs the camera to auto-adjust to best of it's ability.
	// Input can be nil.
	CameraAutoAdjust
)

// Client is the interface should be inherited by all device's pats or modules, it's function is to perform comms with the device.
// Device implementing structure may or may not be a Client, that's for the developer to decide.
//
// invoke - Instruction from a table, to which the client will perform an action.
// data   - The data that the instruction requires. Refer to the documentation of invoke constants, to what that should be
type Client interface {
	Invoke(invoke ClientInvocation, data interface{})
}
