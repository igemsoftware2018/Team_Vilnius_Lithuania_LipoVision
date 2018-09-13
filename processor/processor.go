package processor

import "context"

//Processor Defines a worker for
type Processor interface {

	//Process Performs a process operation on channel items
	Process(queue <-chan struct{}) error

	//WithContext Sets a Context for the process to cancel on
	WithContext(context.Context) Processor
}
