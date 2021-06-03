package controller

import (
	"log"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		log.Println("Received an Options request for user signup")
		return
	}

	// Query db to get all upcoming events: Event[]
	//

	// Convert Event[] to json array

	// Send json response


	log.Println("Received a GET request for retrieving the entire dashboard")
}