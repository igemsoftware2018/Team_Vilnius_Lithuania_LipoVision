package dropletgenomics

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
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
func (p *Pump) Invoke(invoke device.ClientInvocation, data float64) error {
	var (
		endpoint    string
		payloadData map[string][]string
	)

	payloadData = make(map[string][]string)
	payloadData["pump"] = []string{fmt.Sprintf("%d", p.PumpID)}
	payloadData["value"] = []string{fmt.Sprintf("%d", int(data))}

	endpoint = p.BaseAddr + "/update"
	switch invoke {
	case device.PumpSetTargetVolume:
		payloadData["par"] = []string{"volumeTargetW"}
	case device.PumpReset:
	case device.PumpToggleWithdrawInfuse:
		payloadData["par"] = []string{"direction"}
	case device.PumpSetVolume:
		payloadData["par"] = []string{"rate"}
	case device.PumpToggle:
		payloadData["par"] = []string{"status"}
	case device.PumpRefresh:
		payloadData["par"] = []string{"status"}
		endpoint = p.BaseAddr + "/refresh"
	case device.PumpPurge:
		// TODO : collect data in GMC
	default:
		panic("incorrect invoke operation of pump client")
	}

	log.WithFields(log.Fields{"device": "DropletGenomics", "PumpID": p.PumpID, "Parameter": payloadData["par"], "data": data}).Info("Invoke called")

	httpResponse, postErr := MakePost(endpoint, payloadData)
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
