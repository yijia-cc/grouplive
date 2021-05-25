package controller

import (
	"github.com/yijia-cc/grouplive/auth/model"
	"log"
)

func checkAccess(userName string) bool {
	user, err := model.GetUserByName(userName)
	if err != nil {
		return false
	}

	log.Println("[Debug] user in checkAccess(): ", user)

	return user.Role == "admin"
}
