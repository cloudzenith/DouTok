package protobufvalidator

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
)

func Test_doValidate(t *testing.T) {
	req := proto.Message(nil)
	err := doValidate(req)
	assert.Nil(t, err)
}
