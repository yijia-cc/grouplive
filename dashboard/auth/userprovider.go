package auth

import "github.com/yijia-cc/grouplive/dashboard/entity"

type UserProvider interface {
	GetUser(userID string) (entity.User, error)
}
