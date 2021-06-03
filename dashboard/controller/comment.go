package controller

import (
	"log"
	"net/http"
)

// PUT request
func commentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Fetch the comment an event id from the request body

	// Insert the new comment into user_reactions table

	// Send the response status code (success or failure) to the client

	log.Println("Received a PUT request for adding a user comment")
}
