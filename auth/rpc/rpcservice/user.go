package rpcservice

import (
	"context"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/service"
	"github.com/yijia-cc/grouplive/auth/tx"
	"github.com/yijia-cc/grouplive/proto/golang"
)

var _ pb.UserServiceServer = (*User)(nil)

type User struct {
	pb.UnimplementedUserServiceServer
	userService service.User
}

func (u User) GetUser(_ context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	user, err := u.userService.GetUser(request.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        string(user.ID),
		Lastname:  &user.LastName,
		Firstname: &user.FirstName,
		Username: user.Username,
		Unit: &pb.Unit{
			Address:   user.Unit.Address,
			AptNumber: user.Unit.AptNumber,
		},
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}

func NewUser(txFactory tx.TransactionFactory, userDao dao.User) User {
	return User{userService: service.NewUser(txFactory, userDao)}
}
