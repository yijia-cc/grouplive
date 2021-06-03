package controller

import (
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/model"
	"net/http"
)

// AccessHandler implements the http.Handler interface with Access control
type AccessHandler struct{
	access func(string) bool
	handle func(http.ResponseWriter, *http.Request)
}

func (h AccessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userName := getUserNameFromToken(r)

	// User Authorization (Access Control): serve authorized user only
	if h.access(userName) {
		h.handle(w, r)
	} else {
		fmt.Printf("%v is not an authorized user for this resource!!", userName)
		http.Error(w, fmt.Sprintf("%v is not an authorized user for this resource!!", userName), http.StatusForbidden)
	}
}

func checkAccess(userName string) bool {
	user, err := model.GetUserById(userName)
	if err != nil {
		return false
	}

	return user.Role == "admin"
}