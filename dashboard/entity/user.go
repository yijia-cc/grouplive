package entity

import (
	pb "github.com/yijia-cc/grouplive/proto/golang"
	"time"
)

type User struct {
	ID                ID         `json:"id,omitempty"`
	Username          string     `json:"username,omitempty"`
	EncryptedPassword string     `json:"encrypted_password,omitempty"`
	FirstName         string     `json:"first_name,omitempty"`
	LastName          string     `json:"last_name,omitempty"`
	Email             string     `json:"email,omitempty"`
	Phone             string     `json:"phone,omitempty"`
	Unit              Unit       `json:"unit,omitempty"`
	LastSignedInAt    *time.Time `json:"last_signed_in_at,omitempty"`
}

func NewUserFromProto(pbUser *pb.User) User {
	return User{
		ID:        ID(pbUser.GetId()),
		LastName:  pbUser.GetLastname(),
		FirstName: pbUser.GetFirstname(),
		Unit:      NewUnitFromProto(pbUser.Unit),
		Username:  pbUser.GetUsername(),
		Email:     pbUser.GetEmail(),
		Phone:     pbUser.GetPhone(),
	}
}


