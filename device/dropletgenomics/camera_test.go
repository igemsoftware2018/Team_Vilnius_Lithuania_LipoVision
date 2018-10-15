package dropletgenomics_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

func TestAutoAdjust(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"success\": true}")
	}))
	defer ts.Close()

	camera := dropletgenomics.Camera{BaseAddr: ts.URL}

	if err := camera.Invoke(device.CameraAutoAdjust, 0); err != nil {
		t.Fail()
	}
}

func TestSettables(t *testing.T) {
	testCases := []device.ClientInvocation{
		device.CameraSetIllumination,
		device.CameraSetExposure,
		device.CameraSetFrameRate,
	}

	testRetrievers := []func(c dropletgenomics.Camera) int{
		func(c dropletgenomics.Camera) int { return c.Illumination },
		func(c dropletgenomics.Camera) int { return c.Exposure },
		func(c dropletgenomics.Camera) int { return c.FrameRate },
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"success\": true}")
	}))
	defer ts.Close()

	for i, invoke := range testCases {
		value := 10
		camera := dropletgenomics.Camera{BaseAddr: ts.URL}

		if err := camera.Invoke(invoke, value); err != nil {
			t.Error("failed invoke: ", err)
		}

		returned := testRetrievers[i](camera)
		if returned != value {
			t.Error("incorrect return value. got: ", returned)
		}
	}
}
