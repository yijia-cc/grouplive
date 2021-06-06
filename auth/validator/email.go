package validator

import (
	"regexp"
)

var emailFormat = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func ValidateEmail(email *string) bool {
	if email == nil || len(*email) == 0 {
		return true
	}
	return emailFormat.MatchString(*email)
}
