package auth

import "github.com/yijia-cc/grouplive/calendar/entity"

type Authenticator interface {
	GetUser(authToken string) *entity.User
}

var _ Authenticator = (*GroupLiveAuthenticator)(nil)

type GroupLiveAuthenticator struct {
}

func (g GroupLiveAuthenticator) GetUser(authToken string) *entity.User {
	panic("implement me")
}

func NewGroupLiveAuthenticator() GroupLiveAuthenticator {
	return GroupLiveAuthenticator{}
}
