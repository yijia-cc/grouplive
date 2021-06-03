package controller

import (
	"log"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		log.Println("Received an Options request for user signup")
		return
	}

	// Retrieve search criteria from the request body or URL parameters

	// Query db for the specific search criterion

	// Convert data to json

	// Send json response


	log.Println("Received a GET request for search")
}
