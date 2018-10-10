//go:generate mockgen -destination mock_image/mock_image.go image Image
//go:generate mockgen -destination mock_device/mock_device.go github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device Frame
//go:generate mockgen -destination mock_filter/mock_filter.go github.com/Vilnius-Lithuania-iGEM-2018/lipovision/filter Filter
package processor_test

import (
	"context"
	"testing"
	"time"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/processor"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/processor/mock_device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/processor/mock_image"
	"github.com/golang/mock/gomock"
)

func frameGen(ctx context.Context, stream chan device.Frame, getFrame func() device.Frame) {
	go func() {
	FrameGenerate:
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				break FrameGenerate
			default:
				stream <- getFrame()
			}
		}
	}()
}

func TestProcess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	processor := processor.CreateFrameProcessor()
	stream := make(chan device.Frame, 10)
	ctx, cancel := context.WithCancel(context.Background())

	frameGen(ctx, stream, func() device.Frame {
		frameCtx, frameCancel := context.WithTimeout(ctx, 5*time.Millisecond)
		mockImage := mock_image.NewMockImage(mockCtrl)
		mockFrame := mock_device.NewMockFrame(mockCtrl)
		mockFrame.EXPECT().Frame().Return(mockImage).Do(func() {
			frameCancel()
		}).Times(1)
		mockFrame.EXPECT().Skip().Return(frameCtx.Done())
		return mockFrame
	})
	go processor.Process(ctx, stream)

	time.Sleep(50 * time.Millisecond)
	cancel()
}

func TestSkip(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	processor := processor.CreateFrameProcessor()
	stream := make(chan device.Frame, 10)
	ctx, cancel := context.WithCancel(context.Background())

	frameGen(ctx, stream, func() device.Frame {
		frameCtx, frameCancel := context.WithTimeout(ctx, time.Millisecond)
		mockFrame := mock_device.NewMockFrame(mockCtrl)
		mockFrame.EXPECT().Frame().Return(nil).Do(func() {
			frameCancel()
		})
		mockFrame.EXPECT().Skip().Do(func() <-chan struct{} {
			return frameCtx.Done()
		}).AnyTimes()
		return mockFrame
	})

	time.Sleep(20 * time.Millisecond)
	go processor.Process(ctx, stream)
	time.Sleep(10 * time.Millisecond)
	cancel()
}
