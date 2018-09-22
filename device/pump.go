package device

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type dataPack struct {
	DataEscaped string `json:"data_pack"`
	Success     int    `json:"success"`
}

type pumpNameAndValues map[string]pump

type pump struct {
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
	Par   string  `json:"par"`
	Pump  float64 `json:"pump"`
	Value bool    `json:"value"`
}

type requestForResetBody struct {
	Par  interface{} `json:"par"`
	Pump float64     `json:"pump"`
}

type requestForVolumeBody struct {
	Par   string  `json:"par"`
	Pump  float64 `json:"pump"`
	Value float64 `json:"value"`
}

type response struct {
	Success int `json:"success"`
}

type responseReset struct {
	Success int         `json:"success"`
	Data    interface{} `json:"data"`
}

/* Defining main pump functionality */

func NewPump(pumpID int) pump {
	if pumpID > 0 && pumpID < 4 {
		newPump := pump{}
		newPump.PumpID = float64(pumpID)
		return newPump
	}
	newPump := pump{}
	newPump.PumpID = -1
	return newPump

}

func (p *pump) updatePumpValues(updateEndpoint string) bool {

	payload := requestBody{"volume", p.PumpID, true}
	res := makeHTTPRequest(updateEndpoint, &payload)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data dataPack
	var pumpValues pumpNameAndValues
	err = json.Unmarshal(body, &data)
	if isError(err) {
		return false
	}
	if data.Success != 1 {
		return false
	}
	fixedEscape := strings.Replace(data.DataEscaped, "\\", "", -1)
	err = json.Unmarshal([]byte(fixedEscape), &pumpValues)
	if isError(err) {
		return false
	}

	p.VolumeTarget = pumpValues[fmt.Sprint(p.PumpID)].VolumeTarget
	p.PurgeRate = pumpValues[fmt.Sprint(p.PumpID)].PurgeRate
	p.PumpID = pumpValues[fmt.Sprint(p.PumpID)].PumpID
	p.RateW = pumpValues[fmt.Sprint(p.PumpID)].RateW
	p.Volume = pumpValues[fmt.Sprint(p.PumpID)].Volume
	p.Status = pumpValues[fmt.Sprint(p.PumpID)].Status
	p.Name = pumpValues[fmt.Sprint(p.PumpID)].Name
	p.Direction = pumpValues[fmt.Sprint(p.PumpID)].Direction
	p.Syringe = pumpValues[fmt.Sprint(p.PumpID)].Syringe
	p.Used = pumpValues[fmt.Sprint(p.PumpID)].Used
	p.VolumeTargetW = pumpValues[fmt.Sprint(p.PumpID)].VolumeTargetW
	p.VolumeW = pumpValues[fmt.Sprint(p.PumpID)].VolumeW
	p.Rate = pumpValues[fmt.Sprint(p.PumpID)].Rate
	p.Stalled = pumpValues[fmt.Sprint(p.PumpID)].Stalled
	p.Force = pumpValues[fmt.Sprint(p.PumpID)].Force
	p.initialized = true
	return true
}

func (p *pump) togglePump(startEndpoint string, start bool) bool {
	startRequestPayload := requestBody{"status", p.PumpID, start}
	res := makeHTTPRequest(startEndpoint, &startRequestPayload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}
	var responseStruct response
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success == 1 {
		return true
	}
	return false
}

func (p *pump) setVolume(volumeEndpoint string, volume int) bool {
	volumePayload := requestForVolumeBody{"rate", p.PumpID, float64(volume)}
	res := makeHTTPRequest(volumeEndpoint, &volumePayload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}

	var responseStruct response
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success == 1 {
		return true
	}
	return false
}

func (p *pump) toggleWithdrawInfuse(withdrawInfuse string, withdraw bool) bool {
	payload := requestBody{"direction", p.PumpID, withdraw}
	res := makeHTTPRequest(withdrawInfuse, &payload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}

	var responseStruct response
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success == 1 {
		return true
	}
	return false
}

func (p *pump) resetPump(resetEndpoint string) bool {
	payload := requestForResetBody{nil, p.PumpID}
	res := makeHTTPRequest(resetEndpoint, &payload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}

	var responseStruct responseReset
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success == 1 {
		return true
	}
	return false
}

func (p *pump) setTargetVolume(targetEndpoint string, volume int) bool {
	targetVolumePayload := requestForVolumeBody{"volumeTargetW", p.PumpID, float64(volume)}
	res := makeHTTPRequest(targetEndpoint, &targetVolumePayload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}

	var responseStruct response
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success == 1 {
		return true
	}
	return false
}

func (p *pump) purge(purgeEndpoint string) {
	// TODO : collect data in GMC
}
