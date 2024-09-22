package utils

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/proto"
)

func Validate(value proto.Message) error {
	v, err := protovalidate.New()
	if err != nil {
		return err
	}

	if err := v.Validate(value); err != nil {
		return err
	}

	return nil
}
