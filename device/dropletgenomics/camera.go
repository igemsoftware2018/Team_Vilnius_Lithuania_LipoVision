package dropletgenomics

import (
	"encoding/json"
	"errors"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

// Camera has the camera dataset and controlls comms
type Camera struct {
	BaseAddr     string `json:"-"`
	FrameRate    int    `json:"volumeTarget"`
	Exposure     int    `json:"purge_rate"`
	Illumination int    `json:"pump_id"`
}

type payload struct {
	Par   string      `json:"par"`
	Value interface{} `json:"value"`
}

type responseBool struct {
	Success bool `json:"success"`
}

// Invoke performs communications with the device by specific commands
func (c *Camera) Invoke(invoke device.ClientInvocation, data interface{}) error {
	var (
		endpoint    string
		payloadData interface{}
	)

	switch invoke {
	case device.CameraSetExposure:
		endpoint = c.BaseAddr + "/update"
		payloadData = MakePayload("exposure", data)
		c.Exposure = data.(int)
	case device.CameraSetFrameRate:
		endpoint = c.BaseAddr + "/update"
		payloadData = MakePayload("live_rate", data)
		c.FrameRate = data.(int)
	case device.CameraSetIllumination:
		endpoint = c.BaseAddr + "/update"
		payloadData = MakePayload("illumination", data)
		c.Illumination = data.(int)
	case device.CameraAutoAdjust:
		endpoint = c.BaseAddr + "/auto_adjust"
		payloadData = nil
	default:
		panic("incorrect invoke operation of camera client")
	}

	response, postErr := MakePost(endpoint, "application/json", payloadData)
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
