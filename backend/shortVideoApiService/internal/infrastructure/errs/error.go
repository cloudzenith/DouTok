package errs

import "github.com/cloudzenith/DouTok/backend/gopkgs/errorx"

const (
	ErrUnknownUserInfo    = 100001
	ErrFailed2GetUserInfo = 100002
)

func RegisterErrors() {
	errorx.RegisterErrors(ErrUnknownUserInfo, "unknown user info")
	errorx.RegisterErrors(ErrFailed2GetUserInfo, "failed to get user info")
}
