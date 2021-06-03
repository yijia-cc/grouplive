package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Debug]: ", "Received a delete request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Determine user id from token
	userName := getUserNameFromToken(r)

	// Retrieve the event id from the URL path parameter /post/{id}
	id := mux.Vars(r)["id"]

	success, err := deleteEventById(id, userName)
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

// faked for now, will be done soon.
func deleteEventById(id string, userId string) (bool, error) {
	return true, nil
}
