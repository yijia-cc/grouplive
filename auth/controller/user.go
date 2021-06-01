package controller

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/yijia-cc/grouplive/auth/model"
	"net/http"
)

func userIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == "OPTIONS" {
		return
	}

	userId := getUserNameFromToken(r)

	w.Write([]byte(userId))
}


func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	userId := getUserNameFromToken(r)
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading user database %s", err), http.StatusInternalServerError)
	}

	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to serialize user info into JSON %s", err), http.StatusInternalServerError)
	}

	w.Write(js)
}


func getUserNameFromToken(r *http.Request) string {
	userToken := r.Context().Value("userToken")
	claims := userToken.(*jwt.Token).Claims
	return claims.(jwt.MapClaims)["username"].(string)
}
