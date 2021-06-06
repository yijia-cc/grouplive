package validator

import (
	"regexp"
)

var usernameFormat = regexp.MustCompile(`^[0-9a-zA-Z]+$`)

func ValidateUsername(username string) bool {
	return usernameFormat.MatchString(username)
}
