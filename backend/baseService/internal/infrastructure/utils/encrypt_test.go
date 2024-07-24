package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMd5WithSalt(t *testing.T) {
	salt, err := GetPasswordSalt()
	assert.NoError(t, err)
	assert.Equal(t, 32, len(salt))

	md5 := GenerateMd5WithSalt("password", salt)
	md51 := GenerateMd5WithSalt("password", salt)
	assert.Equal(t, md5, md51)
}
