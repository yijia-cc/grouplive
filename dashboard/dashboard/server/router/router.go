package router

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/dashboard/config"
	"github.com/yijia-cc/grouplive/dashboard/service"
	"net/http"
)

func StartupHttpRouter(cfg *config.Config, jwtMiddleware *jwtmiddleware.JWTMiddleware) *mux.Router {
	r := mux.NewRouter()


	// 3 valid route variables are accepted: mixed/grouped/dashboard
	r.Handle("/search/{search_type}/", http.HandlerFunc(service.SearchHandler)).Methods("GET", "OPTIONS")
	r.Handle("/post", http.HandlerFunc(service.PostHandler)).Methods("POST", "OPTIONS")
	r.Handle("/confirm", http.HandlerFunc(service.ConfirmHandler)).Methods("POST", "OPTIONS")

	r.Handle("/update", http.HandlerFunc(service.UpdateHandler)).Methods("PUT", "OPTIONS")

	r.Handle("/delete/{id}", http.HandlerFunc(service.DeleteHandler)).Methods("DELETE", "OPTIONS")

	//r.Handle("/delete/{id}", jwtMiddleware.Handler(http.HandlerFunc(service.DeleteHandler))).Methods("DELETE", "OPTIONS")

	// Serve any files located at <STAIC_MEDIA_DIR>/<filename> via the URL: http://host:port/<STAIC_MEDIA_DIR>/<filename>
	r.PathPrefix(cfg.App.StaticMediaDir[1:]).Handler(http.StripPrefix(cfg.App.StaticMediaDir[1:], http.FileServer(http.Dir(cfg.App.StaticMediaDir))))


	// API that needs both authentication and role-based access control (RBAC)
	//r.Handle("/delete/{id}", jwtMiddleware.Handler(AccessHandler{checkAccess, deleteHandler})).Methods("DELETE", "OPTIONS")

	return r
}