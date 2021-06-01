package controller

import (
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/yijia-cc/grouplive/dashboard/model"
	"log"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		log.Println("Received an Options request for user signup")
		return
	}

	// Determine user id from token
	// userName := getUserNameFromToken(r)

	//resp, err := http.Get("..../auth/userid")  // httpClient.Get()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var userid string
 	//resp.Body.Read(&userid)




	// Query event db for the specific user

	// Convert data to json

	// Send json response
	var event model.Event
	js, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

	log.Println("Received a GET request for search")
}


func getUserNameFromToken(r *http.Request) string {
	userToken := r.Context().Value("userToken")
	claims := userToken.(*jwt.Token).Claims
	return claims.(jwt.MapClaims)["username"].(string)
}