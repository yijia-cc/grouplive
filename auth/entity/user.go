package entity

import (
	"errors"
	"github.com/yijia-cc/grouplive/auth/validator"
	"time"
)

type User struct {
	ID                ID
	LastName          string
	FirstName         string
	Unit              Unit
	Username          string
	Email             string
	EncryptedPassword string
	LastSignedInAt    *time.Time
}

func (u User) Validate() error {
	if !validator.ValidateUsername(u.Username) {
		return errors.New("username is invalid")
	}
	if !validator.ValidateEmail(u.Email) {
		return errors.New("email is invalid")
	}
	return nil
}
