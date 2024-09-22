package utils

import "math/rand"

var numericCharacters = []rune("0123456789")

// GenerateNumericString generate a string only contains numeric characters
func GenerateNumericString(bits int64) string {
	result := make([]rune, bits)
	for i := range result {
		result[i] = numericCharacters[rand.Intn(int(bits))]
	}

	return string(result)
}
