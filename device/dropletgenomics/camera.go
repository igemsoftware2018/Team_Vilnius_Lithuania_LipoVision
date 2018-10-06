package dropletgenomics

import (
	"bytes"
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

func makePayload(setting string, data interface{}) payload {
	return payload{Par: setting, Value: data.(float64)}
}

func makePost(url string, contentType string, data interface{}, response *http.Response) error {
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(&data)
	postResp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return err
	}
	response = postResp
	return nil
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
