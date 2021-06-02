package entity

import "time"

type User struct {
	ID                ID
	Name              *string
	Username          *string
	EncryptedPassword *string
	LastSignedInAt    *time.Time
}
