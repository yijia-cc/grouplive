package controller

import (
	"log"
	"net/http"
)

// PUT request
func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Fetch the event id, and new event data from the request body

	// Update the database for the given event id

	// Send the response status code (success or failure) to the client

	log.Println("Received a PUT request for updating an existing event")
}
