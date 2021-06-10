package auth

import (
	"github.com/yijia-cc/grouplive/dashboard/auth/permission"
	"github.com/yijia-cc/grouplive/dashboard/entity"
)

type Authorizer interface {
	HasPermission(user *entity.User, permission permission.Permission) bool
}
