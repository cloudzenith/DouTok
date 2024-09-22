package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// GetPasswordSalt returns a random salt string with the given length
func GetPasswordSalt() (string, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func GenerateMd5WithSalt(password, salt string) string {
	password += salt
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
