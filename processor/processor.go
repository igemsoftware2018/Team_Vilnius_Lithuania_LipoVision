package processor

// Processor Defines a worker for
type Processor interface {
	// Process Performs a process operation on channel items
	Process(<-chan struct{})
}
