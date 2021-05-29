package controller

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/yijia-cc/grouplive/auth/model"
	"log"
	"net/http"
	"time"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Debug]:", "Received a Sign-in request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == "OPTIONS" {
		return
	}

	// Get User information from client sign in form
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Unable to decode user data from client: %v\n", err), http.StatusBadRequest)
		log.Printf("Unable to decode user data from client: %v\n", err)
		return
	}

	fmt.Println(user)

	// Check if a specified username and password exist and match each other
	if err := model.Auth(&user); err != nil {
		http.Error(w, "User doesn't exists or wrong password", http.StatusUnauthorized)
		fmt.Printf("User doesn't exists or wrong password\n")
		return
	}

	fmt.Println("[Debug]:", user, "logged in successfully!")

	// Create a token with specified Claims (Payload), Payload = {"username": xxx, "exp": yyy}, and encoding algorithm (HS256).
	// The actual encoding is not done until SigningKey is provided in token.SignedString(SigningKey).
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Set the token expiration time to 24 hours from now
	})

	// Token encoding and signing: get the complete, signed token using our provided secret key. The actual token encoding
	// is done once we provide the key here.
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		log.Printf("Failed to generate token %v\n", err)
		return
	}

	// Send the token back to client in text/plain response
	w.Write([]byte(tokenString))
}