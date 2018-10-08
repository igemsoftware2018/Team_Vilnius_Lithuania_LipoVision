package dropletgenomics_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

func TestNonSettables(t *testing.T) {
	testCases := []device.ClientInvocation{
		device.PumpSetTargetVolume,
		device.PumpReset,
		device.PumpToggleWithdrawInfuse,
		device.PumpSetVolume,
		device.PumpToggle,
		device.PumpPurge,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"success\": true}")
	}))
	defer ts.Close()

	for _, invoke := range testCases {
		value := 10
		pump := dropletgenomics.Pump{BaseAddr: ts.URL}

		if err := pump.Invoke(invoke, value); err != nil {
			t.Error("failed invoke: ", err)
		}
	}
}

func TestRefresh(t *testing.T) {
	doubleJson := `
	{
		"volumeTarget": 123.1,
		"purge_rate": 654.1,
		"pump_id": 132,
		"rateW": 189.1,
		"volume": 85.8,
		"status": true,
		"name": "testing_pump",
		"direction": false,
		"syringe": 123.56,
		"used": true,
		"volumeTargetW": 1459.4,
		"volumeW": 1774.4,
		"rate": 12987.4,
		"stalled": true,
		"force": 289621.4
	}
	`
	refreshJson := `
	{
		"data_pack": "%s",
		"success": true
	}
	`

	escapedJson := strings.Replace(strings.Replace(doubleJson, "\n", "", -1), "\"", "\\\"", -1)
	payload := fmt.Sprintf(strings.Replace(refreshJson, "\n", "", -1), escapedJson)
	payload = strings.Replace(payload, "\t", "", -1)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, payload)
	}))
	defer ts.Close()

	pump := dropletgenomics.Pump{BaseAddr: ts.URL}
	if err := pump.Invoke(device.PumpRefresh, nil); err != nil {
		t.Error("invoke failed: ", err)
		return
	}

	equlitySlice := []bool{
		pump.VolumeTarget == 123.1,
		pump.PurgeRate == 654.1,
		pump.PumpID == 132,
		pump.RateW == 189.1,
		pump.Volume == 85.8,
		pump.Status,
		strings.Compare(pump.Name, "testing_pump") == 0,
		!pump.Direction,
		pump.Syringe == 123.56,
		pump.Used,
		pump.VolumeTargetW == 1459.4,
		pump.VolumeW == 1774.4,
		pump.Rate == 12987.4,
		pump.Stalled,
		pump.Force == 289621.4,
	}

	for _, equality := range equlitySlice {
		if !equality {
			t.Error("some equality check failed")
			return
		}
	}
}
