package dropletgenomics

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// MakePayload creates a payload structure
// TODO: eliminate use
func MakePayload(setting string, data interface{}) payload {
	return payload{Par: setting, Value: data}
}

// MakePost encodes a struct to json and sends a post to url
// TODO: eliminate use
func MakePost(url string, contentType string, data interface{}) (*http.Response, error) {
	reqBody := new(bytes.Buffer)

	err := json.NewEncoder(reqBody).Encode(data)
	if err != nil {
		return nil, err
	}

	postResp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return nil, err
	}

	return postResp, nil
}
