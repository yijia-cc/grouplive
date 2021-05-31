package auth

import (
	"github.com/yijia-cc/grouplive/calendar/auth/permission"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type Authorizer interface {
	HasPermission(user *entity.User, permission permission.Permission, resourceID *string) bool
}