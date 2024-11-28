package errorx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(1, "some error")
	assert.NotNil(t, err)
	assert.Equal(t, int32(1), err.Code)
	assert.Equal(t, "some error", err.Error())
}

func TestNewWithCode(t *testing.T) {
	err := NewWithCode(1)
	assert.NotNil(t, err)
	assert.Equal(t, int32(1), err.Code)
	assert.Equal(t, "unknown error", err.Error())

	RegisterErrors(1, "some error")
	err = NewWithCode(1)
	assert.NotNil(t, err)
	assert.Equal(t, int32(1), err.Code)
	assert.Equal(t, "some error", err.Error())
}
