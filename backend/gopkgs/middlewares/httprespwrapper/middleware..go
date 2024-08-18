package httprespwrapper

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/gopkgs/errorx"
	"github.com/go-kratos/kratos/v2/middleware"
)

type Wrapper struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func HttpResponseWrapper() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			result, err := handler(ctx, req)
			response := &Wrapper{}
			if err == nil {
				response.Code = errorx.SuccessCode
				response.Msg = errorx.SuccessMsg
				response.Data = result
				return response, nil
			}

			var errx *errorx.Error
			ok := errors.As(err, &errx)
			if !ok {
				response.Code = errorx.UnknownErrorCode
				response.Msg = err.Error()
				return response, nil
			}

			response.Code = errx.Code
			response.Msg = errx.Msg
			return response, nil
		}
	}
}
