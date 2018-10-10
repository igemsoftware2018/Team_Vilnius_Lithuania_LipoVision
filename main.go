package main

import (
	"flag"
	"fmt"
)

func main() {
	deviceRequested := flag.String("device", "",
		`Specify a device to run the program with. Valid values:
	 	 > dropletgenomics
		  > video`)
	flag.Parse()

	//processor := processor.CreateFrameProcessor()
	fmt.Printf(*deviceRequested)
}
