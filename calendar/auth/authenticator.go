package auth

import "github.com/yijia-cc/grouplive/calendar/entity"

type Authenticator struct {

}

func (a Authenticator) GetUser (authToken string) entity.User {
	panic("Implement me!")
}