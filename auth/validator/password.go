package validator

func ValidatePassword(password string) bool {
	return len(password) >= 8
}
