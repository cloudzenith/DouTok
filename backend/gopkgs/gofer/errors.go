package gofer

var (
	ErrSubmittedTaskNil     = "submitted task can't be nil"
	ErrGroupHashBeenStarted = "can't submit task to a error group after Wait()"
)
