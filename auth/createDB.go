package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yijia-cc/grouplive/auth/config"
	"github.com/yijia-cc/grouplive/auth/model"
)

func main() {
	pool := model.DBConn()
	defer pool.Close()

	log.Printf("Connected to DB %s successfully!\n", config.UserDBConfig.DBname)

	if err := model.CreateUserTable(); err != nil {
		log.Printf("Error %s when creating users table", err)
		return
	}

	admin := model.User{UserName: "isabella", Password: "1234", FirstName: "Qingqing", LastName: "Kang", Email: "isabellakqq@gmail.com", Apt: "123", Role: "admin"}
	if res, err := model.AddUser(&admin); err != nil {
		log.Printf("Error %s when inserting user %s\n", err, admin.UserName)
	} else {
		affected, err := res.RowsAffected()
		if err == nil && affected > 0 {
			log.Printf("User %s is added successfully!\n", admin.UserName)
		}
	}

	resident := model.User{UserName: "john123", Password: "1234", FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com", Apt: "234", Role: "resident"}
	if res, err := model.AddUser(&resident); err != nil {
		log.Printf("Error %s when inserting user %s\n", err, resident.UserName)
	} else {
		affected, err := res.RowsAffected()
		if err == nil && affected > 0 {
			log.Printf("User %s is added successfully!\n", resident.UserName)
		}
	}

	log.Printf("Connection to DB <%s> is still alive!\n", config.UserDBConfig.DBname)
}

