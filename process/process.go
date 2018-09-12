package process

//Context Contains all the global variables of a process
type Context struct {
}

//Process Defines a cancellable process
type Process interface {
	Run(Context, *bool) error
}
