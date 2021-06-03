package service

import "github.com/yijia-cc/grouplive/auth/entity"

type User struct {
}

func (User) GetUser(userID string) (entity.User, error) {
	panic("Isabella, implements me!")
}

func NewUser( /*Isabella, inject dependencies here*/ ) User {
	return User{}
}
