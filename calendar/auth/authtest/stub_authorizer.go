package authtest

import (
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/auth/permission"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

var _ auth.Authorizer = (*StubAuthorizer)(nil)

type StubAuthorizer struct {
}

func (s StubAuthorizer) HasPermission(user *entity.User, permission permission.Permission, resourceID *string) bool {
	return true
}

func NewStubAuthorizer() StubAuthorizer {
	return StubAuthorizer{}
}
