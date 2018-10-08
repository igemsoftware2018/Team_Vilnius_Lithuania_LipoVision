package dropletgenomics

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	PumpSetTargetVolume clientInvocation = iota
	PumpReset
	PumpToggleWithdrawInfuse
	PumpSetVolume
	PumpToggle
	PumpRefresh
	PumpPurge
)

type dataPack struct {
	DataEscaped string `json:"data_pack"`
	Success     int    `json:"success"`
}

//Pump owns data of a pump and performs comms with device
type Pump struct {
	VolumeTarget  float64 `json:"volumeTarget"`
	PurgeRate     float64 `json:"purge_rate"`
	PumpID        float64 `json:"pump_id"`
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
	Success int         `json:"success"`
	Data    interface{} `json:"data"`
}

//Invoke performs communications with the device by specific commands
func (p Pump) Invoke(invoke clientInvocation, data interface{}) error {
	const pumpBaseAddr = "http://192.168.1.100:8764"
	var (
		endpoint    string
		payloadData interface{}
	)

	endpoint = pumpBaseAddr + "/update"
	switch invoke {
	case PumpSetTargetVolume:
		payloadData = requestBody{Par: "volumeTargetW", Pump: p.PumpID, Value: data}
	case PumpReset:
		payloadData = requestBody{Pump: p.PumpID}
	case PumpToggleWithdrawInfuse:
		payloadData = requestBody{Par: "direction", Pump: p.PumpID, Value: data}
	case PumpSetVolume:
		payloadData = requestBody{Par: "rate", Pump: p.PumpID, Value: data}
	case PumpToggle:
		payloadData = requestBody{Par: "status", Pump: p.PumpID, Value: data}
	case PumpRefresh:
		payloadData = requestBody{Par: "status", Pump: p.PumpID, Value: data}
		endpoint = pumpBaseAddr + "/refresh"
	case PumpPurge:
		// TODO : collect data in GMC
	default:
		panic("incorrect invoke operation of pump client")
	}

	fmt.Printf("%f\n", p.PumpID)
	httpResponse, postErr := makePost(endpoint, "application/json", payloadData)
	if postErr != nil {
		return postErr
	}

	var responseData response
	switch invoke {
	case PumpRefresh:
		var doubleJSON dataPack
		if err := json.NewDecoder(httpResponse.Body).Decode(&doubleJSON); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(doubleJSON.DataEscaped), &p); err != nil {
			return err
		}
	default:
		if err := json.NewDecoder(httpResponse.Body).Decode(&responseData); err != nil {
			return err
		}
		if responseData.Success == 1 {
			return errors.New("camera device failed to process the request")
		}
	}
	return nil
}
