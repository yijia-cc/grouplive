package rpcservice

import (
	"context"

	"github.com/yijia-cc/grouplive/auth/rpc/proto"
	"github.com/yijia-cc/grouplive/auth/service"
)

var _ proto.UserServiceServer = (*User)(nil)

type User struct {
	proto.UnimplementedUserServiceServer
	userService service.User
}

func (u User) GetUser(_ context.Context, request *proto.GetUserRequest) (*proto.User, error) {
	user, err := u.userService.GetUser(request.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.User{
		Id:        string(user.ID),
		Lastname:  &user.LastName,
		Firstname: &user.FirstName,
		Unit: &proto.Unit{
			Address:   user.Unit.Address,
			AptNumber: user.Unit.AptNumber,
		},
	}, nil
}

func NewUser() User {
	return User{userService: service.NewUser()}
}
