package service

import (
	"encoding/json"
	"github.com/yijia-cc/grouplive/dashboard/auth"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"log"
	"net/http"
)

// UserHandler handles a GET request to return the meta data of user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	log.Println("Received a GET request for retrieving user info/confirmation")

	// Determine search type from URL
	st, err := getSearchType(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	var js []byte

	user, err := auth.UserFromContext(r.Context())
	if err != nil {
		http.Error(w, "missing user info in request context", http.StatusInternalServerError)
		log.Println("missing user info in request context")
		return
	}

	switch st {
	case entity.UserInfo:
		js, err = json.Marshal(user)
	case entity.UserConfirmation:
		js, err = dao.NewReactionDao(db.DashDB).GetConfirmations(user.Username)
	}

	// Query db for the specific search criterion
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	log.Println("successful")
}