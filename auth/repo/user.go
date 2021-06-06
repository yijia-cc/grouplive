package repo

import (
	"errors"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type User struct {
	userDao dao.User
}

type FindUserQuery struct {
	Username *string `json:"username"`
	ID       *string `json:"user_id"`
	Email    *string `json:"user_email"`
}

func (u User) FindUser(tx tx.Transaction, query FindUserQuery) (entity.User, error) {
	if query.ID != nil {
		return u.userDao.FindUserByID(tx, query.ID)
	}
	if query.Username != nil {
		return u.userDao.FindUserByUsername(tx, query.Username)
	}
	if query.Email != nil {
		return u.userDao.FindUserByEmail(tx, query.Email)
	}
	return entity.User{}, errors.New("user filtering condition not provided")
}

func (u User) CreateUser(tx tx.Transaction, user entity.User) error {
	return u.userDao.CreateUser(tx, user)
}

func NewUser(userDao dao.User) User {
	return User{userDao: userDao}
}
