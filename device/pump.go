package device

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type dataPack struct {
	DataEscaped string `json:"data_pack"`
	Success     int    `json:"success"`
}

type pumpNameAndValues map[string]pump

type pump struct {
	VolumeTarget  float64 `json:"volumeTarget"`
	PurgeRate     int     `json:"purge_rate"`
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
	Force         int     `json:"force"`
	initialized   bool    `json:"-"`
}

type requestBody struct {
	Par   string `json:"par"`
	Pump  int    `json:"pump"`
	Value bool   `json:"value"`
}

type response struct {
	Success int `json:"success"`
}

/* Defining main pump functionality */

func (p *pump) updatePumpValues(updateEndpoint string) bool {

	updateRequestBodyPayload := requestBody{"volume", p.PumpID, false}
	res := makeHTTPRequest(updateEndpoint, &updateRequestBodyPayload)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data dataPack
	var pumpValues pumpNameAndValues
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&data)
	if isError(err) {
		return false
	}
	if data.Success != 1 {
		return false
	}
	fixedEscapeString := strings.Replace(string(data.DataEscaped), "\\", "", -1)
	err = json.NewDecoder(strings.NewReader(fixedEscapeString)).Decode(&pumpValues)
	if isError(err) {
		return false
	}

	p.VolumeTarget = pumpValues[strconv.Itoa(p.PumpID)].VolumeTarget
	p.PurgeRate = pumpValues[strconv.Itoa(p.PumpID)].PurgeRate
	p.PumpID = pumpValues[strconv.Itoa(p.PumpID)].PumpID
	p.RateW = pumpValues[strconv.Itoa(p.PumpID)].RateW
	p.Volume = pumpValues[strconv.Itoa(p.PumpID)].Volume
	p.Status = pumpValues[strconv.Itoa(p.PumpID)].Status
	p.Name = pumpValues[strconv.Itoa(p.PumpID)].Name
	p.Direction = pumpValues[strconv.Itoa(p.PumpID)].Direction
	p.Syringe = pumpValues[strconv.Itoa(p.PumpID)].Syringe
	p.Used = pumpValues[strconv.Itoa(p.PumpID)].Used
	p.VolumeTargetW = pumpValues[strconv.Itoa(p.PumpID)].VolumeTargetW
	p.VolumeW = pumpValues[strconv.Itoa(p.PumpID)].VolumeW
	p.Rate = pumpValues[strconv.Itoa(p.PumpID)].Rate
	p.Stalled = pumpValues[strconv.Itoa(p.PumpID)].Stalled
	p.Force = pumpValues[strconv.Itoa(p.PumpID)].Force
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

func (p *pump) purge(purgeEndpoint string) {
	// TODO
}

// func (p *pump) setVolume(volumeEndpoint string, volume int) bool {
// 	volumePayload := requestBody{"rate", p.PumpID, volume}
// 	res := makeHTTPRequest(volumeEndpoint, &startRequestPayload)

// 	responseBody, _ := ioutil.ReadAll(res.Body)
// 	if responseBody == nil {
// 		return false
// 	}

// 	var responseStruct response
// 	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
// 	if isError(err) {
// 		return false
// 	}
// 	if responseStruct.Success == 1 {
// 		return true
// 	}
// 	return false
// }

/* Helper functions */
func makeHTTPRequest(endpointURL string, sendBody interface{}) *http.Response {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&sendBody)
	req, _ := http.NewRequest("POST", endpointURL, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if isError(err) {
		res = nil
	}
	if res.StatusCode != 200 {
		res = nil
	}
	return res
}

func isError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
