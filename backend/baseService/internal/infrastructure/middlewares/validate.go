package middlewares

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/protobuf/proto"
)

func ProtobufValidator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			value, ok := req.(proto.Message)
			if !ok {
				return nil, errors.New("not valid request")
			}

			v, err := protovalidate.New()
			if err != nil {
				return nil, err
			}

			if err := v.Validate(value); err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}
