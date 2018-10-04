package dropletgenomics

import (
	"encoding/json"
	"io/ioutil"
	"strings"
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

func (c *camera) setIllumination(endpoint string, illumination float64) bool {
	c.Illumination = illumination

	ilPayload := payload{"illumination", illumination}
	res := makeHTTPRequest(endpoint, &ilPayload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}
	var responseStruct responseBool
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success {
		return true
	}
	return false
}

func (c *camera) setExposure(endpoint string, exposure float64) bool {
	c.Exposure = exposure

	exPayload := payload{"exposure", exposure}
	res := makeHTTPRequest(endpoint, &exPayload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}
	var responseStruct responseBool
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success {
		return true
	}
	return false
}

func (c *camera) setFrameRate(endpoint string, frameRate float64) bool {
	c.FrameRate = frameRate

	frPayload := payload{"live_rate", frameRate}
	res := makeHTTPRequest(endpoint, &frPayload)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}
	var responseStruct responseBool
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success {
		return true
	}
	return false
}

func (c *camera) autoAdjust(endpoint string) bool {
	res := makeHTTPRequest(endpoint, nil)

	responseBody, _ := ioutil.ReadAll(res.Body)
	if responseBody == nil {
		return false
	}
	var responseStruct responseBool
	err := json.NewDecoder(strings.NewReader(string(responseBody))).Decode(&responseStruct)
	if isError(err) {
		return false
	}
	if responseStruct.Success {
		return true
	}
	return false
}
