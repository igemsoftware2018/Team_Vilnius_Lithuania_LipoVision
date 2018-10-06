package dropletgenomics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type clientInvocation int

//client is the interface that is inherited by all device's pats or modules, whose data is received non-locally
type client interface {
	Invoke(clientInvocation, interface{})
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

/* Helper functions */
func makeHTTPRequest(endpointURL string, sendBody interface{}) *http.Response {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&sendBody)
	req, _ := http.NewRequest("POST", endpointURL, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if isError(err) {
		res = nil
		return res
	}
	if res.StatusCode != 200 {
		res = nil
		return res
	}
	return res
}

func isError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func setupDeviceURL(device *Device) string {
	return fmt.Sprintf("http://%v:%v", device.IPAddress, strconv.Itoa(device.HTTPPort))
}
func (device *Device) setupGetValuesURL() string {
	return fmt.Sprintf("http://%v:%v/refresh", device.IPAddress, strconv.Itoa(device.PumpDataPort))
}
func (device *Device) setupToggleURL() string {
	return fmt.Sprintf("http://%v:%v/update", device.IPAddress, strconv.Itoa(device.PumpDataPort))
}
func formatPumpValues(p *Pump) string {
	return fmt.Sprintf("Pump: %.0f\n\tVolume Target: %.0f\n\tPurge Rate: %.0f\n\tRateW: %.0f\n\tVolume: %.0f\n\tStatus: %v\n\tName: %s\n\tDirection: %v\n\tSyringe: %.2f\n\tUsed: %v\n\tVolumeTargetW: %.0f\n\tVolumeW: %.0f\n\tRate: %.0f\n\tStalled: %v\n\tForce: %.0f\n", p.PumpID, p.VolumeTarget, p.PurgeRate, p.RateW, p.Volume, p.Status, p.Name, p.Direction, p.Syringe, p.Used, p.VolumeTargetW, p.VolumeW, p.Rate, p.Stalled, p.Force)
}
