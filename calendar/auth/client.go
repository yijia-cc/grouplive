package auth

import (
	"github.com/yijia-cc/grouplive/calendar/auth/permission"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

var _ Authenticator = (*Client)(nil)
var _ Authorizer = (*Client)(nil)

type Client struct {
}

func (c Client) HasPermission(user *entity.User, permission permission.Permission, resourceID *string) bool {
	panic("implement me")
}

func (c Client) GetUser(authToken string) entity.User {
	panic("implement me")
}

func NewClient() Client{
	return Client{}
}

