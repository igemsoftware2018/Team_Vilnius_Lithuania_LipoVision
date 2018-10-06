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

func makePayload(setting string, data interface{}) payload {
	return payload{Par: setting, Value: data.(float64)}
}

func (c camera) Invoke(invoke clientInvocation, data interface{}) error {
	const cameraBaseAddr string = "http://192.168.1.100:8765"

	var endpoint string
	var payloadData payload

	switch invoke {
	case CameraSetExposure:
		endpoint = cameraBaseAddr + "/update"
	case CameraSetFrameRate:
		endpoint = cameraBaseAddr + "/update"
	case CameraSetIllumination:
		endpoint = cameraBaseAddr + "/update"
		payloadData = makePayload("illumination", data)
	case CameraAutoAdjust:
		endpoint = cameraBaseAddr + "/auto_adjust"
	default:
		panic("incorrect invoke operation of camera client")
	}

	response := makeHTTPRequest(endpoint, payloadData)
	if response != nil {
		return errors.New("failed to communicate with camera")
	}

	var responseData responseBool
	err := json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return err
	}

	if !responseData.Success {
		return errors.New("camera device failed to process the request")
	}

	return nil
}
