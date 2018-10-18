package video

import (
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
)

// NewPump returns a fake pump
func NewPump() Pump {
	return Pump{}
}

// Pump is merely a mock object for actual pumps
// camera's don't actually have pumps
type Pump struct {
	device.Client
}

// Invoke on a camera's pump does not do anything, it's
// a camera, it does not have pumps
func (Pump) Invoke(invoke device.ClientInvocation, data float64) error {
	log.WithFields(log.Fields{
		"device": "video",
	}).Info("Invoke on pump called")
	return nil
}
