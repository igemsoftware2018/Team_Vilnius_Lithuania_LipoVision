package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	device := dropletgenomics.Create()

	fmt.Printf("%s\n", "Waiting for device")
	for !device.Available() {
		time.Sleep(2 * time.Second)
	}

	if device.Available() {
		fmt.Printf("%s\n", "Connected!")
	}

	for {
		select {
		case <-ctx.Done():
			break
		default:
			if err := device.RefreshAll(); err != nil {
				fmt.Printf("%s\n", err)
				time.Sleep(time.Second)
			}
			for i := 0; i < device.NumPumps(); i++ {
				fmt.Printf("Pump %d values: %f\n", i, device.Pump(i).Rate)
			}
		}
	}

	cancel()
}
