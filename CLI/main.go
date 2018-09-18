package main

import (
	"flag"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

func main() {

	var numberOfPumpsUsed int
	var ipAddress = flag.String("ip", "localhost", "Enter IP address of Device")
	var httpPort = flag.Int("hp", 5000, "http port of the device")
	var pumpDataPort = flag.Int("dp", 5000, "Data port of the device")
	var pumpCameraPort = flag.Int("cp", 5000, "Camera port of the device")
	flag.Parse()

	deviceInstance := device.DropletGenomicsDevice{*ipAddress, *httpPort, *pumpDataPort, *pumpCameraPort, 0, nil}

	if deviceInstance.Available() {
		fmt.Print("Device availabe, select how many pumps do you want in experiment:\n")
		_, err := fmt.Scanf("%d", &numberOfPumpsUsed)
		if err != nil {
			panic("Wrong input")
		}
		deviceInstance.PumpExperiment = numberOfPumpsUsed
		deviceInstance.EstablishPumps()
		if deviceInstance.Update() {
			fmt.Printf("%v", "Updating pumps succeeded")
		}
		fmt.Printf("%v", numberOfPumpsUsed)
	} else {
		fmt.Print("Device is not availabe! Shutting down...")
	}
}
