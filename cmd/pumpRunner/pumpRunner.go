package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	device := dropletgenomics.Create()

	signalBuffer := make(chan os.Signal, 1)
	signal.Notify(signalBuffer, os.Interrupt)
	go func() {
		for sig := range signalBuffer {
			fmt.Printf("signal: %s\n", sig)
			cancel()
			break
		}
	}()

	fmt.Printf("%s\n", "Waiting for device")
	for !device.Available() {
		time.Sleep(2 * time.Second)
	}

	if device.Available() {
		fmt.Printf("%s\n", "Connected!")
	}

Processing:
	for {
		select {
		case <-ctx.Done():
			break Processing
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
}
