package errorx

import "sync"

const (
	SuccessCode      = 0
	SuccessMsg       = "success"
	UnknownErrorCode = -1
	UnknownErrorMsg  = "unknown error"
)

var globalErrorCode = sync.Map{}

func RegisterErrors(code int32, msg string) {
	globalErrorCode.Store(code, msg)
}

type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int32, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func NewWithCode(code int32) *Error {
	msg, ok := globalErrorCode.Load(code)
	if !ok {
		msg = UnknownErrorMsg
	}

	return &Error{
		Code: code,
		Msg:  msg.(string),
	}
}

func (e *Error) Error() string {
	return e.Msg
}
