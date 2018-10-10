package processor

import (
	"context"
)

// Processor Defines a worker for
type Processor interface {
	// Process Performs a process operation on channel items
	Process(context.Context, <-chan struct{})
}
