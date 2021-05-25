package controller

import (
	"encoding/json"
	"fmt"
	"github.com/yijia-cc/grouplive/auth/model"
	"log"
	"net/http"
	"regexp"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow CORS only for any origins
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == "OPTIONS" {
		//log.Println("Received an Options request for user signup")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Unable to decode user data from client: %v\n", err), http.StatusBadRequest)
		log.Printf("Unable to decode user data from client: %v\n", err)
		return
	}

	// regexp ^[a-z0-9]$ means every character must be a-z or 0-9. ^ means the begining, $ means the end.
	if user.UserName == "" || user.Password == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.UserName) {
		http.Error(w, "Invalid user id or password", http.StatusBadRequest)
		fmt.Printf("Invalid username or password\n")
		return
	}

	// Set default user role as "resident"
	if user.Role == "" {
		user.Role = "resident"
	}

	if res, err := model.AddUser(&user); err != nil {
		log.Printf("Error %s when inserting user %s\n", err, user.UserName)
		http.Error(w, "User already exists", http.StatusBadRequest)
	} else {
		affected, err := res.RowsAffected()
		if err == nil && affected > 0 {
			log.Printf("User %s is registered successfully!\n", user.UserName)
		}
	}
}
