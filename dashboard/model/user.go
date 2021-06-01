package model

type User struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Apt       string `json:"apt"`
	Role      string `json:"role"`
}


func GetUserById(userName string) (*User, error) {
	var user User
	sql := "SELECT username, password, first_name, last_name, email, apt, role from users where username = ?"

	err := userDB.QueryRow(sql, userName).Scan(&user.UserName, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.Apt, &user.Role)
	return &user, err
}