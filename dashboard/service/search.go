package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/db/dao"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"log"
	"net/http"
)

// SearchHandler handles a GET request to search the event using a given search criterion
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	log.Println("Received a GET request for searching the event")

	// Determine search type from URL
	searchType, err := getSearchType(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	fmt.Println(searchType)

	// Parse search keys for search type == mixed/grouped
	searchKeys := entity.SearchKey{}
	if searchType != entity.Dashboard {
		if err := json.NewDecoder(r.Body).Decode(&searchKeys); err != nil {
			http.Error(w, fmt.Sprintf("unable to decode search keys from request body: %s", err), http.StatusBadRequest)
			log.Printf("unable to decode search keys from request body: %s\n", err)
			return
		}
	}

	//fmt.Println(searchKeys)

	// Query db for the specific search criterion
	js, err := dao.NewEventDao(db.DashDB).Search(searchType, &searchKeys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	log.Println("Event search successful")
}

func getSearchType(r *http.Request) (entity.SearchType, error) {
	st := mux.Vars(r)["search_type"]
	switch st {
	case "mixed":
		return entity.Mixed, nil
	case "grouped":
		return entity.Grouped, nil
	case "dashboard":
		return entity.Dashboard, nil
	default:
		return -1, fmt.Errorf("invalid search type: %s", st)
	}
}