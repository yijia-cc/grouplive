package service

import (
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"log"
	"net/http"
)

// MetaHandler handles a GET request to return the meta data of event table
func MetaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	log.Println("Received a GET request for retrieving meta data")

	// Query db for the specific search criterion
	js, err := dao.NewEventDao(db.DashDB).MetaMapping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	log.Println("Meta data obtained successfully")
}


