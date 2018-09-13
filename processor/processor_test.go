package processor

import (
	"context"
	"testing"
	"time"
)

func createIntProcessor() intProcessor {
	return intProcessor{ctx: context.Background()}
}

type intProcessor struct {
	ctx context.Context
}

func (ip intProcessor) Process(queue <-chan int) error {
	for item := range queue {
		select {
		case <-ip.ctx.Done():
			return
		default:
		}

		// Do something
	}
}

func (ip intProcessor) WithContext(ctx context.Context) intProcessor {
	ip.ctx = ctx
	return ip
}

func TestProcessesItems(*testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	processor := createIntProcessor().WithContext(ctx)
	stream := make(chan int, 5)

	go func() {
		for i := 0; i < 20; i++ {
			stream <- i
			time.Sleep(5 * time.Millisecond)
		}
	}()

	go processor.Process(stream)

	time.Sleep(time.Second)
	cancel()
}
