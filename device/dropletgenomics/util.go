package dropletgenomics

import (
	"net/http"
)

// MakePayload creates a payload structure
// TODO: eliminate use
func MakePayload(setting string, data interface{}) payload {
	return payload{Par: setting, Value: data}
}

// MakePost encodes a struct to json and sends a post to url
// TODO: eliminate use
func MakePost(url string, data map[string][]string) (*http.Response, error) {

	postResp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}

	return postResp, nil
}
