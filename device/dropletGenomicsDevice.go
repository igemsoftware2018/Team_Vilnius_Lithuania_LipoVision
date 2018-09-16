package device

import (
	"fmt"
	"net/http"
	"strconv"
)

// DropletGenomicsDevice implements default Device interface
type DropletGenomicsDevice struct {
	ipAddress         string
	httpPort          int
	pumpDataPort      int
	recondingDataPort int
	pumpExperiment    int
	pumps             []pump
}

// Available determines whether connection to
// DropletGenomics device is established
func (device *DropletGenomicsDevice) Available() bool {
	url := setupDeviceURL(device)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false
	}
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return false
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return false
	}

	return true
}

func setupDeviceURL(device *DropletGenomicsDevice) string {
	return fmt.Sprintf("http://%v:%v", device.ipAddress, strconv.Itoa(device.httpPort))
}
