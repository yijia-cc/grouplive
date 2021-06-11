package authtest

import (
	"github.com/yijia-cc/grouplive/dashboard/auth"
	"github.com/yijia-cc/grouplive/dashboard/auth/permission"
	"github.com/yijia-cc/grouplive/dashboard/entity"
)

var _ auth.Authorizer = (*StubAuthorizer)(nil)

type StubAuthorizer struct {
}

func (s StubAuthorizer) HasPermission(user *entity.User, permission permission.Permission) bool {
	return true
}

func NewStubAuthorizer() StubAuthorizer {
	return StubAuthorizer{}
}
