package main

import (
	"flag"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/window"
)

func main() {
	deviceRequested := flag.String("device", "",
		`Specify a device to run the program with. Valid values:
	 	 > dropletgenomics
		  > video`)
	flag.Parse()

	mainWindow, err := window.NewMain()
	if err != nil {
		panic(err)
	}
	mainWindow.Run()

	//processor := processor.CreateFrameProcessor()
	fmt.Printf(*deviceRequested)
}
