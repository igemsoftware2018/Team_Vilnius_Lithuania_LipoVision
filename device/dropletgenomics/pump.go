package dropletgenomics

import (
	"encoding/json"
	"errors"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

type dataPack struct {
	DataEscaped string `json:"data_pack"`
	Success     bool   `json:"success"`
}

// Pump owns data of a pump and performs comms with device
type Pump struct {
	BaseAddr      string  `json:"-"`
	VolumeTarget  float64 `json:"volumeTarget"`
	PurgeRate     float64 `json:"purge_rate"`
	PumpID        int     `json:"pump_id"`
	RateW         float64 `json:"rateW"`
	Volume        float64 `json:"volume"`
	Status        bool    `json:"status"`
	Name          string  `json:"name"`
	Direction     bool    `json:"direction"`
	Syringe       float64 `json:"syringe"`
	Used          bool    `json:"used"`
	VolumeTargetW float64 `json:"volumeTargetW"`
	VolumeW       float64 `json:"volumeW"`
	Rate          float64 `json:"rate"`
	Stalled       bool    `json:"stalled"`
	Force         float64 `json:"force"`
	initialized   bool
}

type requestBody struct {
	Par   interface{} `json:"par"`
	Pump  interface{} `json:"pump"`
	Value interface{} `json:"value"`
}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Invoke performs communications with the device by specific commands
func (p *Pump) Invoke(invoke device.ClientInvocation, data interface{}) error {
	var (
		endpoint    string
		payloadData interface{}
	)

	endpoint = p.BaseAddr + "/update"
	switch invoke {
	case device.PumpSetTargetVolume:
		payloadData = requestBody{Par: "volumeTargetW", Pump: p.PumpID, Value: data}
	case device.PumpReset:
		payloadData = requestBody{Pump: p.PumpID}
	case device.PumpToggleWithdrawInfuse:
		payloadData = requestBody{Par: "direction", Pump: p.PumpID, Value: data}
	case device.PumpSetVolume:
		payloadData = requestBody{Par: "rate", Pump: p.PumpID, Value: data}
	case device.PumpToggle:
		payloadData = requestBody{Par: "status", Pump: p.PumpID, Value: data}
	case device.PumpRefresh:
		payloadData = requestBody{Par: "status", Pump: p.PumpID, Value: data}
		endpoint = p.BaseAddr + "/refresh"
	case device.PumpPurge:
		// TODO : collect data in GMC
	default:
		panic("incorrect invoke operation of pump client")
	}

	httpResponse, postErr := MakePost(endpoint, "application/json", payloadData)
	if postErr != nil {
		return postErr
	}

	var decodeError error
	switch invoke {
	case device.PumpRefresh:
		var doubleJSON dataPack
		if decodeError = json.NewDecoder(httpResponse.Body).Decode(&doubleJSON); decodeError == nil {
			if decodeError := json.Unmarshal([]byte(doubleJSON.DataEscaped), &p); decodeError == nil {
				return nil
			}
		}
	default:
		var responseData response
		if decodeError = json.NewDecoder(httpResponse.Body).Decode(&responseData); decodeError == nil {
			if !responseData.Success {
				decodeError = errors.New("camera device failed to process the request")
			}
		}
	}
	return decodeError
}
