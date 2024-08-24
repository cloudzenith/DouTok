package respcheck

import (
	"errors"
)

type metadata interface {
	GetBizCode() int32
	GetMessage() string
	GetDomain() string
	GetReason() []string
}

type response[T metadata] interface {
	GetMeta() T
}

func Check[Meta metadata](resp response[Meta], e error) error {
	if e != nil {
		return e
	}

	if resp == nil {
		return errors.New("response is nil")
	}

	meta := resp.GetMeta()
	if meta.GetBizCode() == 0 {
		return nil
	}

	if meta.GetMessage() != "" {
		return errors.New(meta.GetMessage())
	}

	return errors.New("unknown error")
}

func CheckT[T any, Meta metadata](resp response[Meta], e error, noError func() T) (t T, err error) {
	if e != nil {
		return t, e
	}

	if resp == nil {
		return t, errors.New("response is nil")
	}

	meta := resp.GetMeta()
	if meta.GetBizCode() == 0 {
		t = noError()
		return t, nil
	}

	if meta.GetMessage() != "" {
		return t, errors.New(meta.GetMessage())
	}

	return t, errors.New("unknown error")
}
