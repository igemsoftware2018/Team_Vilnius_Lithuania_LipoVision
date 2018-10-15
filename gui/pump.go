package gui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// NewPumpAndRegionContainer creates a box container for
// both RegionControl and PumpControl
func NewPumpAndRegionContainer() (gtk.IWidget, *RegionControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if boxErr != nil {
		return nil, nil, boxErr
	}

	region, regionErr := NewRegionControl()
	if regionErr != nil {
		return nil, nil, regionErr
	}

	pumpControl, pumpControlErr := NewPumpControl()
	if pumpControlErr != nil {
		return nil, nil, pumpControlErr
	}

	box.PackStart(pumpControl.Root(), true, true, 0)
	box.PackEnd(region.Root(), false, true, 0)

	return box, region, nil
}

// NewPumpControl returns pump control collection
func NewPumpControl() (*PumpControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if boxErr != nil {
		return nil, boxErr
	}

	frame, frameErr := gtk.FrameNew("Pump controls")
	if frameErr != nil {
		return nil, frameErr
	}

	var pumps []gtk.IWidget
	for i := 0; i < 4; i++ {
		pump, err := newPumpItem(i + 1)
		if err != nil {
			return nil, err
		}
		pumps = append(pumps, pump)
		box.PackStart(pump, false, false, 5)
	}

	frame.Add(box)

	return &PumpControl{rootBox: frame, pumps: pumps}, nil
}

func newPumpItem(index int) (gtk.IWidget, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if boxErr != nil {
		return nil, boxErr
	}

	label, labelErr := gtk.LabelNew(fmt.Sprintf("Pump %d: ", index))
	if labelErr != nil {
		return nil, labelErr
	}

	spinbox, spinboxErr := gtk.SpinButtonNewWithRange(-9000, 9000, 5)
	if spinboxErr != nil {
		return nil, spinboxErr
	}
	spinbox.SetValue(0)

	box.PackStart(label, false, false, 0)
	box.PackStart(spinbox, true, true, 0)

	return box, nil
}

// PumpControl is the pump control collection
type PumpControl struct {
	Control

	rootBox *gtk.Frame
	pumps   []gtk.IWidget
}

// Root returns the root stack component
func (pc *PumpControl) Root() gtk.IWidget {
	return pc.rootBox
}
