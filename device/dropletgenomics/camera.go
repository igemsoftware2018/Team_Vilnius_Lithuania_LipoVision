package dropletgenomics

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Comms operations for the camera
const (
	CameraSetIllumination clientInvocation = iota
	CameraSetExposure
	CameraSetFrameRate
	CameraAutoAdjust
)

type camera struct {
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

func (c camera) Invoke(invoke clientInvocation, data interface{}) error {
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

	var response *http.Response
	if err := makePost(endpoint, "application/json", payloadData, response); err != nil {
		return err
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
