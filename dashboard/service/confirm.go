package service

import (
	"encoding/json"
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"log"
	"net/http"
	"time"
)

// ConfirmHandler handles a POST request to confirm either a reservation or a cancellation of an event
func ConfirmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	log.Println("Received a POST request for confirming/unconfirming an event")

	now := time.Now()
	re := entity.Reaction{
 		CreatedAt: now,
 		UpdatedAt: now,
 		Active: true,
	}

	// Retrieve the user re from request body
	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to decode user re: %s", err), http.StatusBadRequest)
		log.Printf("unable to decode user re: %s", err)
		return
	}

	reDao := dao.NewReactionDao(db.DashDB)
	rid, err := reDao.GetReactionId(re.Username, re.EventId)
	if err != nil {
		http.Error(w, fmt.Sprintf("database failure while querying the Reaction table: %s", err), http.StatusInternalServerError)
		return
	}

	// confirm or unconfirm
	if rid <= 0 {
		_, err = reDao.Creat(&re)
	} else {
		//_, err = reDao.ToggleById(rid)
		re.Id = rid
		_, err = reDao.Update(&re)
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to confirm/unconfirm: %v", err), http.StatusInternalServerError)
		log.Println("failed to confirm/unconfirm:", err)
		return
	}

	log.Println("Successfully confirmed/unconfirmed")
}