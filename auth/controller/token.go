package controller

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"log"
	"net/http"
)

// tokenHandler is for tesing purpose only. It displays all key/value pairs stored in the user token
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow CORS only for any origins
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == "OPTIONS" {
		//log.Println("Received an Options request for user signup")
		return
	}

	// Retrieve decoded user token from the request context
	userToken := r.Context().Value("userToken")

	fmt.Println("UserToken = ", userToken)

	claims := userToken.(*jwt.Token).Claims
	userName := claims.(jwt.MapClaims)["username"].(string)
	log.Println("[Debug] userName = ", userName)

	for k, v := range claims.(jwt.MapClaims) {
		fmt.Fprintf(w, "%s: %#v\n", k, v)
	}
}
