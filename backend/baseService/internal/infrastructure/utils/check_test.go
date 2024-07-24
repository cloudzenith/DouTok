package utils

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestIsValidWithRegex(t *testing.T) {
	t.Run("test mobile", func(t *testing.T) {
		ok := IsValidWithRegex(`^[0-9]*$`, "1234567890")
		assert.Equal(t, ok, true)
	})
}
