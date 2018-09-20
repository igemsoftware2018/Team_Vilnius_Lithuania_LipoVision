package main

import (
	"flag"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

/* CLI - initializing device object and manipulate its methods */

func main() {

	var numberOfPumpsUsed int
	var userCLIDialogSelection int
	var ipAddress = flag.String("ip", "localhost", "Enter IP address of Device")
	var httpPort = flag.Int("hp", 5000, "http port of the device")
	var pumpDataPort = flag.Int("dp", 5000, "Data port of the device")
	var pumpCameraPort = flag.Int("cp", 5000, "Camera port of the device")
	flag.Parse()

	deviceInstance := device.DropletGenomicsDevice{IPAddress: *ipAddress, HTTPPort: *httpPort, PumpDataPort: *pumpDataPort, RecondingDataPort: *pumpCameraPort, PumpExperiment: 0, Pumps: nil}

	if deviceInstance.Available() {
		fmt.Print("Device availabe, select how many pumps do you want in experiment:\n")
		_, err := fmt.Scanf("%d\n", &numberOfPumpsUsed)
		if err != nil {
			panic("Wrong input")
		}
		deviceInstance.PumpExperiment = numberOfPumpsUsed
		deviceInstance.EstablishPumpsWithId()
		if deviceInstance.Update() {
			active := true
			for active {
				fmt.Print(getSelections())
				_, err := fmt.Scanf("%d\n", &userCLIDialogSelection)
				if err != nil {
					fmt.Println(err)
				}
				switch userCLIDialogSelection {
				case 0:
					checkDeviceStatus(&deviceInstance)
				case 1:
					getPumpValues(&deviceInstance)
				case 2:
					togglePump(&deviceInstance)
				case -1:
					active = false
				default:
					continue
				}
			}
			fmt.Println("Exiting program...")
		}
	} else {
		fmt.Print("Device is not availabe! Shutting down...\n")
	}
}

func getPumpValues(device *device.DropletGenomicsDevice) {
	if device.Available() {
		var selectedPump int
		fmt.Println("Select pump (-1 for all)")
		fmt.Scanf("%d\n", &selectedPump)
		if selectedPump == -1 {
			fmt.Print(device.GetPumpValues(selectedPump))
		} else {
			fmt.Print(device.GetPumpValues(selectedPump))
		}
	}
}
func checkDeviceStatus(device *device.DropletGenomicsDevice) {
	if device.Available() {
		fmt.Println("\nDevice is Available")
	} else {
		fmt.Println("\nDevice is not available")
	}
}
func togglePump(device *device.DropletGenomicsDevice) {
	if device.Available() {

	}
}

func getSelections() string {
	return (`
Please enter one of the following:

0 - check device status
1 - get pump values
2 - toggle pump
3 - purge
-1 - exit program

`)
}
