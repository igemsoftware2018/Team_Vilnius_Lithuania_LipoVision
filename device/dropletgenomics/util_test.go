package dropletgenomics_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

type TestStr struct {
	Success string `json:"success"`
}

func TestPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resBytes, _ := ioutil.ReadAll(r.Body)
		text := string(resBytes[:])
		fmt.Fprintln(w, text)
	}))
	defer ts.Close()

	testVal := TestStr{Success: "yes"}
	response, err := dropletgenomics.MakePost(ts.URL, "text/html", testVal)
	if err != nil {
		t.Error("failed to make post with: ", err)
		return
	}

	var result TestStr
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Log(result)
		t.Error(err)
	}

	if strings.Compare(testVal.Success, result.Success) != 0 {
		t.Error("mismatch, got: ", result.Success)
	}
}
