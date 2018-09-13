package processor

import "context"

//Processor Defines a worker for
type Processor interface {

	//Process Performs a process operation on channel items
	Process(<-chan struct{})

	//WithContext Sets a Context for the process to cancel on
	WithContext(context.Context) Processor
}
