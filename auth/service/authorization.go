package service

import "github.com/yijia-cc/grouplive/auth/entity"

type Authorization struct {
}

func (Authorization) HasPermission(permissionID string, userID *entity.ID, resourceID *entity.ID) bool {
	panic("Isabella, implements me!")
}

func NewAuthorization( /*Isabella, inject dependencies here*/ ) Authorization {
	return Authorization{}
}
