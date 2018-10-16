package dropletgenomics

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
)

// Camera has the camera dataset and controlls comms
type Camera struct {
	BaseAddr     string  `json:"-"`
	FrameRate    float64 `json:"volumeTarget"`
	Exposure     float64 `json:"purge_rate"`
	Illumination float64 `json:"pump_id"`
}

type payload struct {
	Par   string      `json:"par"`
	Value interface{} `json:"value"`
}

type responseBool struct {
	Success bool `json:"success"`
}

func makePayload(parameter string, data float64) map[string][]string {
	payloadData := make(map[string][]string)
	payloadData["par"] = []string{parameter}
	payloadData["value"] = []string{fmt.Sprintf("%d", int(data))}
	return payloadData
}

// Invoke performs communications with the device by specific commands
func (c *Camera) Invoke(invoke device.ClientInvocation, data float64) error {
	var (
		endpoint    string
		payloadData map[string][]string
	)

	switch invoke {
	case device.CameraSetExposure:
		endpoint = c.BaseAddr + "/update"
		payloadData = makePayload("exposure", data)
		c.Exposure = data
		log.WithFields(log.Fields{
			"device": "DropletGenomics",
			"value":  data,
		}).Info("Exposure set")
	case device.CameraSetFrameRate:
		endpoint = c.BaseAddr + "/update"
		payloadData = makePayload("live_rate", data)
		c.FrameRate = data
		log.WithFields(log.Fields{
			"device": "DropletGenomics",
			"value":  data,
		}).Info("Framerate set")
	case device.CameraSetIllumination:
		endpoint = c.BaseAddr + "/update"
		payloadData = makePayload("illumination", data)
		c.Illumination = data
		log.WithFields(log.Fields{
			"device": "DropletGenomics",
			"value":  data,
		}).Info("Illumination set")
	case device.CameraAutoAdjust:
		endpoint = c.BaseAddr + "/auto_adjust"
		payloadData = nil
		log.WithFields(log.Fields{
			"device": "DropletGenomics",
		}).Info("Auto-adjust called")
	default:
		panic("incorrect invoke operation of camera client")
	}

	response, postErr := MakePost(endpoint, payloadData)
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
