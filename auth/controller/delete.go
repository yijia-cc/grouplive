package controller

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Debug]: ", "Received a delete request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	// Retrieve decoded user token from the request context
	userToken := r.Context().Value("userToken")
	claims := userToken.(*jwt.Token).Claims
	userName := claims.(jwt.MapClaims)["username"].(string)

	// Retrieve the URL path parameter from the URL /post/{id} matched by the mux HTTP router
	id := mux.Vars(r)["id"]

	//fmt.Printf("Delete post: id = %s, user = %s\n", id, username)

	success, err := deletePost(id, userName)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete post from Elasticsearch: %v\n", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	if success {
		//w.Write([]byte("Post is deleted successfully"))
		//fmt.Println("Post is deleted successfully")

		w.Write([]byte(fmt.Sprintf("%s is an admin who has access to this resource", userName)))
		fmt.Printf("%s is an admin who has access to this resource", userName)
	}
}

// faked right now for demo purpose!
func deletePost(id string, userId string) (bool, error) {
	return true, nil
}