package service

import (
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/auth"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"github.com/yijia-cc/grouplive/dashboard/util"
	"log"
	"net/http"
	"time"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	log.Println("Received a POST request for creating a new event")

	user, err := auth.UserFromContext(r.Context())
	if err != nil {
		http.Error(w, "missing user info in request context", http.StatusInternalServerError)
		log.Println("missing user info in request context")
		return
	}

	event, mediaList, status, err := util.ParseMultipartRequest(r)
	if err != nil {
		http.Error(w, err.Error(), status)
		log.Println(err.Error())
		return
	}

	event.User = user
	event.CreatedAt = time.Now()
	event.UpdatedAt = event.CreatedAt

	err = createEvent(event, mediaList)
	if err != nil {
		http.Error(w, fmt.Sprintf("server failed to save event to database: %v", err), http.StatusInternalServerError)
		log.Println("server failed to save event to database:", err)
		return
	}

	log.Println("New event created successfully")

	/*fmt.Printf("Uploaded files: %v\nRecieved form values: %#v\n", mediaList, event)
	fmt.Printf("Start time: %s, end time: %s\n", event.StartTime, event.EndTime)
	fmt.Println("Formated start time:", event.StartTime.Format(LOCAL_DATE_TIME_FMT))
	fmt.Printf("Local start time: %d-%d-%d %d-%d-%d\n", event.StartTime.Year(), event.StartTime.Month(), event.StartTime.Day(),
		event.StartTime.Hour(), event.StartTime.Minute(), event.StartTime.Second())*/
}

// createEvent save the event and media files using transaction
func createEvent(event *entity.Event, mediaList []*entity.Media) error {
	tx, err := db.DashDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	event.Id, err = dao.NewEventDao(db.DashDB).CreatEventTx(tx, event)
	if err != nil {
		return err
	}

	for _, media := range mediaList {
		media.Event = event
		media.CreatedAt = time.Now()
		media.UpdatedAt = media.CreatedAt
		_, err = dao.NewMediaDao(db.DashDB).CreatMediaTx(tx, media)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}




