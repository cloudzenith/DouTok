package protobufvalidator

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/protobuf/proto"
)

func doValidate(req proto.Message) error {
	v, err := protovalidate.New()
	if err != nil {
		return err
	}

	if err := v.Validate(req); err != nil {
		return err
	}

	return nil
}

func CheckEnv() error {
	_, err := protovalidate.New()
	if err != nil {
		return err
	}

	return nil
}

func ProtobufValidator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			value, ok := req.(proto.Message)
			if !ok {
				return nil, errors.New("not valid request")
			}

			if err := doValidate(value); err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}
