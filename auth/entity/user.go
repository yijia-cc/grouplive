package entity

import "time"

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
