package controller

import (
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Fetch new event data from request body

	// Save new event to database

	// Send the response status code (success or failure) to the client


	log.Println("Received a Post request for adding a new event")
}
