package dropletgenomics

import (
	"encoding/json"
	"errors"
)

// Comms operations for the camera
const (
	CameraSetIllumination clientInvocation = iota
	CameraSetExposure
	CameraSetFrameRate
	CameraAutoAdjust
)

//Camera has the camera dataset and controlls comms
type Camera struct {
	FrameRate    float64 `json:"volumeTarget"`
	Exposure     float64 `json:"purge_rate"`
	Illumination float64 `json:"pump_id"`
}

type payload struct {
	Par   string  `json:"par"`
	Value float64 `json:"value"`
}

type responseBool struct {
	Success bool `json:"success"`
}

//Invoke performs communications with the device by specific commands
func (c Camera) Invoke(invoke clientInvocation, data interface{}) error {
	const cameraBaseAddr string = "http://192.168.1.100:8765"

	var (
		endpoint    string
		payloadData interface{}
	)

	switch invoke {
	case CameraSetExposure:
		endpoint = cameraBaseAddr + "/update"
		payloadData = makePayload("exposure", data)
	case CameraSetFrameRate:
		endpoint = cameraBaseAddr + "/update"
		payloadData = makePayload("live_rate", data)
	case CameraSetIllumination:
		endpoint = cameraBaseAddr + "/update"
		payloadData = makePayload("illumination", data)
	case CameraAutoAdjust:
		endpoint = cameraBaseAddr + "/auto_adjust"
		payloadData = nil
	default:
		panic("incorrect invoke operation of camera client")
	}

	response, postErr := makePost(endpoint, "application/json", payloadData)
	if postErr != nil {
		return postErr
	}

	var responseData responseBool
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return err
	}

	if !responseData.Success {
		return errors.New("camera device failed to process the request")
	}

	return nil
}
