package errorx

const (
	SuccessCode      = 0
	SuccessMsg       = "success"
	UnknownErrorCode = -1
	UnknownErrorMsg  = "unknown error"
)

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

func (e *Error) Error() string {
	return e.Msg
}
