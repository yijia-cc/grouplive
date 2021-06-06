package service

import (
	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/repo"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type User struct {
	txFactory tx.TransactionFactory
	userRepo  repo.User
}

func (u User) GetUser(userID string) (entity.User, error) {
	query := repo.FindUserQuery{ID: &userID}
	transaction, err := u.txFactory.NewTransaction()
	if err != nil {
		return entity.User{}, err
	}
	return u.userRepo.FindUser(transaction, query)
}

func NewUser(txFactory tx.TransactionFactory, userDao dao.User) User {
	return User{
		txFactory: txFactory,
		userRepo:  repo.NewUser(userDao),
	}
}
