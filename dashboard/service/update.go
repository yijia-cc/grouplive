package service

import (
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"github.com/yijia-cc/grouplive/dashboard/util"
	"log"
	"net/http"
	"time"
)

// UpdateHandler handles a PUT request to update an existing event
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	log.Println("Received a PUT request for updating an existing event")

	// Fetch the event id, and new event data from the request body
	event, mediaList, status, err := util.ParseMultipartRequest(r)
	if err != nil {
		http.Error(w, err.Error(), status)
		log.Println(err.Error())
		return
	}

	event.UpdatedAt = time.Now()

	// Update the database for the given event id
	rowsUpdated, err := updateEvent(event, mediaList)
	if err != nil {
		http.Error(w, fmt.Sprintf("server failed to update the event in DB: %v", err), http.StatusInternalServerError)
		log.Println("failed to update event in database:", err)
		return
	}

	if rowsUpdated == 0 {
		http.Error(w, fmt.Sprintf("unable to update possibly due to nonexistent event id (%d): %v", event.Id, err), http.StatusBadRequest)
		log.Printf("unable to update possibly due to nonexistent event id (%d): %v\n", event.Id, err)
		return
	}

	log.Println("Event updated successfully")
}

// Update the event and media files using transaction
func updateEvent(event *entity.Event, mediaList []*entity.Media) (int64, error) {
	tx, err := db.DashDB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rowsUpdated, err := dao.NewEventDao(db.DashDB).UpdateEventTx(tx, event)
	if err != nil {
		return 0, err
	}

	/*for _, media := range mediaList {
		media.Event = event
		media.UpdatedAt = time.Now()
		_, err = dao.NewMediaDao(db.DashDB).UpdateMediaTx(tx, media)
		if err != nil {
			return err
		}
	}*/

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}