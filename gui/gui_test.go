package gui_test

import (
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/gui"
)

func TestGUILoads(t *testing.T) {
	_, err := gui.NewMainControl()
	if err != nil {
		t.Error("Something failed while assembling components: ", err)
	}
}
