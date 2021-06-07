package model

import "database/sql"

type User struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Apt       string `json:"apt"`
	Role      string `json:"role"`
}


func AddUser(user *User) (sql.Result, error) {
	sql := "INSERT INTO users (username, password, first_name, last_name, email, apt, role) VALUES (?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(sql, user.UserName, user.Password, user.FirstName, user.LastName, user.Email, user.Apt, user.Role)
	return res, err
}


func Auth(user *User) error {
	sql := "SELECT first_name, last_name, email, apt, role from users where username = ? and password = ?"

	// QueryRow always returns a non-nil value (*Row), containing at most one row from the db. If the query selects no rows,
	// Errors are deferred until Row's Scan() method is called, in which case, the (*Row).Scan() method returns a ErrNoRows (error object).
	err := db.QueryRow(sql, user.UserName, user.Password).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Apt, &user.Role)
	return err
}


func GetUserById(userId string) (*User, error) {
	var user User
	sql := "SELECT username, password, first_name, last_name, email, apt, role from users where username = ?"

	err := db.QueryRow(sql, userId).Scan(&user.UserName, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.Apt, &user.Role)
	return &user, err
}
