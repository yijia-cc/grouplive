package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"log"
	"net/http"
	"strconv"
)

// DeleteHandler handles a DELETE request to delete an existing event
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Debug]: ", "Received a delete request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Retrieve the event id from the URL path parameter /post/{id}
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, fmt.Sprintf("invalid id: %v", id), http.StatusBadRequest)
		log.Printf("invalid id encountered: %v", id)
		return
	}

	rowsDeleted, err := deleteEvent(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("server failed to delete: %v", err), http.StatusInternalServerError)
		log.Printf("failed to delete: %v", err)
		return
	}

	if rowsDeleted == 0 {
		http.Error(w, fmt.Sprintf("unable to delete possibly due to nonexistent event id: %d", id), http.StatusBadRequest)
		log.Printf("unable to delete possibly due to nonexistent event id: %d", id)
		return
	}

	log.Println("Event deleted successfully")
}

// deleteEvent delete event and the associated media files for a given event id
func deleteEvent(eid int64) (int64, error) {
	tx, err := db.DashDB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	eventDao := dao.NewEventDao(db.DashDB)
	mediaDao := dao.NewMediaDao(db.DashDB)
	reactionDao := dao.NewReactionDao(db.DashDB)

	rowsAffected, err := eventDao.DeleteByEidTx(tx, eid)
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, nil
	}

	_, err = mediaDao.DeleteByEidTx(tx, eid)
	if err != nil {
		return 0, err
	}

	_, err = reactionDao.DeleteByEidTx(tx, eid)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}