package dao

import (
	"database/sql"

	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type PermissionBinding interface {
	CountPermissionBindings(tx tx.Transaction, binding entity.PermissionBinding) (int, error)
}

var _ PermissionBinding = (*PermissionBindingSQL)(nil)

type PermissionBindingSQL struct {
	db *sql.DB
}

func (p PermissionBindingSQL) CountPermissionBindings(tx tx.Transaction, binding entity.PermissionBinding) (int, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT COUNT(*)
FROM permission_binding
WHERE permission_id = ? AND user_id = ? AND resource_type = ? AND resource_id = ?;
`, binding.Permission.ID, binding.User.ID, binding.Resource.Type.ID, binding.Resource.ID)

	var count int
	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return 0, NotFound{}
	}
	return count, err
}

func NewPermissionBindingSQL(db *sql.DB) PermissionBindingSQL {
	return PermissionBindingSQL{db: db}
}
