package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Debug]: ", "Received a delete request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Retrieve decoded user token from the request context
	userName := getUserNameFromToken(r)

	// Retrieve the URL path parameter from the URL /post/{id} matched by the mux HTTP routing
	id := mux.Vars(r)["id"]

	success, err := deletePost(id, userName)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete post from Elasticsearch: %v\n", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	if success {
		w.Write([]byte(fmt.Sprintf("%s is an admin who has access to this resource", userName)))
		fmt.Printf("%s is an admin who has access to this resource", userName)
	}
}

// faked right now for demo purpose!
func deletePost(id string, userId string) (bool, error) {
	return true, nil
}
