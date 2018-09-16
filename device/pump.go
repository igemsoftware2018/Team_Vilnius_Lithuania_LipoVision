package device

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
}

type refreshRequestBody struct {
	Par  string `json:"par"`
	Pump int    `json:"pump"`
}

type response struct {
	Success string `json:"success"`
}

func (p *pump) updatePumpValues(updateEndpoint string, specificPumpID int) bool {

	refresh := refreshRequestBody{"volume", specificPumpID}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(refresh)
	req, _ := http.NewRequest("POST", updateEndpoint, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	check(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("POST failed")
	}
	body, _ := ioutil.ReadAll(res.Body)
	var data dataPack
	var pumpValues pumpNameAndValues
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&data)
	check(err)
	if data.Success != 1 {
		return false
		panic("Success value not 1")
	}
	fixed := strings.Replace(string(data.DataEscaped), "\\", "", -1)
	err = json.NewDecoder(strings.NewReader(fixed)).Decode(&pumpValues)

	check(err)
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
