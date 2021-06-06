package service

import (
	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type Authorization struct {
	txFactory            tx.TransactionFactory
	permissionBindingDao dao.PermissionBinding
}

func (a Authorization) HasPermission(permission entity.Permission, user entity.User, resource entity.Resource) (bool, error) {
	transaction, err := a.txFactory.NewTransaction()
	if err != nil {
		return false, err
	}

	binding := entity.PermissionBinding{
		Permission: permission,
		User:       user,
		Resource:   resource,
	}
	count, err := a.permissionBindingDao.CountPermissionBindings(transaction, binding)
	switch err.(type) {
	case dao.NotFound:
		return false, nil
	case nil:
		if count > 0 {
			return true, nil
		} else {
			return false, nil
		}
	default:
		return false, err
	}
}

func NewAuthorization(txFactory tx.TransactionFactory, permissionBindingDao dao.PermissionBinding) Authorization {
	return Authorization{
		txFactory:            txFactory,
		permissionBindingDao: permissionBindingDao,
	}
}
