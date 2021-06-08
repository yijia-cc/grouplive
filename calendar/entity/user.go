package entity

import (
	pb "github.com/yijia-cc/grouplive/proto/golang"
	"time"
)

type User struct {
	ID                ID
	LastName          string
	FirstName         string
	Unit              Unit
	Username          string
	Email             string
	Phone             string
	EncryptedPassword string
	LastSignedInAt    *time.Time
}

func NewUserFromProto(pbUser *pb.User) User {
	return User {
		ID: ID(pbUser.GetId()),
		LastName: pbUser.GetLastname(),
		FirstName: pbUser.GetFirstname(),
		Unit : NewUnitFromProto(pbUser.Unit),
		Username: pbUser.GetUsername(),
		Email: pbUser.GetEmail(),
		Phone: pbUser.GetPhone(),
	}
}
