package auth

import "github.com/yijia-cc/grouplive/calendar/entity"

type Authenticator interface {
	GetUser (authToken string) entity.User
}