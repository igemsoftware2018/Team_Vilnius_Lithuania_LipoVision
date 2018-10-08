package dropletgenomics

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type clientInvocation int

//Client is the interface that is inherited by all device's pats or modules, it's functions perform comms with the device
type Client interface {
	Invoke(clientInvocation, interface{})
}

func makePayload(setting string, data interface{}) payload {
	return payload{Par: setting, Value: data.(float64)}
}

func makePost(url string, contentType string, data interface{}) (*http.Response, error) {
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(data)
	postResp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return nil, err
	}
	return postResp, nil
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
