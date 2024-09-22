package utils

import "regexp"

func IsValidWithRegex(pattern, str string) bool {
	regex := regexp.MustCompile(pattern)
	matched := regex.MatchString(str)
	return matched
}
