package main

import (
	"context"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
	"gocv.io/x/gocv"
)

func main() {
	dropletDevice := dropletgenomics.CreateDropletGenomicsDevice()
	originalWindow := gocv.NewWindow("Stream")
	ctx, cancel := context.WithCancel(context.Background())
	stream := dropletDevice.Stream(ctx)

	for {
		select {
		case frame := <-stream:
			mat, err := gocv.ImageToMatRGB(frame.Frame())
			if err != nil {
				fmt.Println(err)
			}
			originalWindow.IMShow(mat)
			if originalWindow.WaitKey(10)&0xFF == 'c' {
				cancel()
				break
			}
			fmt.Printf("frame done\n")
		}
	}
}
