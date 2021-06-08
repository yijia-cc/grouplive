package auth

import "github.com/yijia-cc/grouplive/calendar/entity"

type UserProvider interface {
	GetUser(userID string) (entity.User, error)
}

