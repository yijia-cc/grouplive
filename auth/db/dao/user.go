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
	CreateUser(tx tx.Transaction, user entity.User) error
}

var _ User = (*UserSQL)(nil)

type UserSQL struct {
	db *sql.DB
}

func (u UserSQL) CreateUser(tx tx.Transaction, user entity.User) error {
	_, err := tx.DBTransaction.Exec(`
INSERT INTO user
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
`, user.ID, user.LastName, user.FirstName, user.Username, user.Email, user.EncryptedPassword, nil, user.Unit.Address, user.Unit.AptNumber)
	return err
}

func (u UserSQL) FindUserByID(tx tx.Transaction, id *string) (entity.User, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT id, last_name, first_name, username, encrypted_password, last_signed_in_at, address, apt_number
FROM user
WHERE id = ?;
`, id)
	return findUser(row)
}

func (u UserSQL) FindUserByUsername(tx tx.Transaction, username *string) (entity.User, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT id, last_name, first_name, username, encrypted_password, last_signed_in_at, address, apt_number
FROM user
WHERE username = ?;
`, username)
	return findUser(row)
}

func (u UserSQL) FindUserByEmail(tx tx.Transaction, email *string) (entity.User, error) {
	row := tx.DBTransaction.QueryRow(`
SELECT id, last_name, first_name, username, encrypted_password, last_signed_in_at, address, apt_number
FROM user
WHERE email = ?;
`, email)
	return findUser(row)
}

func findUser(row *sql.Row) (entity.User, error) {
	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.LastName,
		&user.FirstName,
		&user.Username,
		&user.EncryptedPassword,
		&user.LastSignedInAt,
		&user.Unit.Address,
		&user.Unit.AptNumber,
	)
	if err == sql.ErrNoRows {
		return entity.User{}, NotFound{}
	}
	return user, err
}

func NewUserSQL(db *sql.DB) UserSQL {
	return UserSQL{db: db}
}
