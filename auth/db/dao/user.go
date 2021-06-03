package dao

import (
	"database/sql"

	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type User interface {
	FindUserByID(tx tx.Transaction, id *string) (entity.User, error)
	FindUserByUsername(tx tx.Transaction, username *string) (entity.User, error)
	FindUserByEmail(tx tx.Transaction, email *string) (entity.User, error)
}

var _ User = (*UserSQL)(nil)

type UserSQL struct {
	db *sql.DB
}

func (u UserSQL) FindUserByID(tx tx.Transaction, id *string) (entity.User, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT id, name, username, encrypted_password, last_signed_in_at
FROM user
WHERE id = ?;
`, id)
	return findUser(row)
}

func (u UserSQL) FindUserByUsername(tx tx.Transaction, username *string) (entity.User, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT id, name, username, encrypted_password, last_signed_in_at
FROM user
WHERE username = ?;
`, username)
	return findUser(row)
}

func (u UserSQL) FindUserByEmail(tx tx.Transaction, email *string) (entity.User, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT id, name, username, encrypted_password, last_signed_in_at
FROM user
WHERE email = ?;
`, email)
	return findUser(row)
}

func findUser(row *sql.Row) (entity.User, error) {
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.EncryptedPassword, &user.LastSignedInAt)
	return user, err
}

func NewUserSQL(db *sql.DB) UserSQL {
	return UserSQL{db: db}
}
