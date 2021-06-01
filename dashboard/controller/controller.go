package controller

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/dashboard/config"
	"net/http"
)

var SigningKey []byte

func StartUp(cfg *config.Config) *mux.Router {
	SigningKey = []byte(cfg.TokenSecretKey)
	jwtMiddleware := StartupJWT()
	router := StartupHttpRouter(jwtMiddleware)
	return router
}

func StartupJWT() *jwtmiddleware.JWTMiddleware {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return SigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		UserProperty: "userToken",
	})

	return jwtMiddleware
}

func StartupHttpRouter(jwtMiddleware *jwtmiddleware.JWTMiddleware) *mux.Router {
	r := mux.NewRouter()

	// API that needs user authentication ONLY
	r.Handle("/user", jwtMiddleware.Handler(http.HandlerFunc(userHandler))).Methods("GET", "OPTIONS")
	//r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS")
	//r.Handle("/dashboard", jwtMiddleware.Handler(http.HandlerFunc(dashboardHandler))).Methods("GET", "OPTIONS")
	//
	//r.Handle("/post", jwtMiddleware.Handler(http.HandlerFunc(postHandler))).Methods("POST", "OPTIONS")
	//
	//r.Handle("/update", jwtMiddleware.Handler(http.HandlerFunc(updateHandler))).Methods("PUT", "OPTIONS")
	//r.Handle("/comment", jwtMiddleware.Handler(http.HandlerFunc(commentHandler))).Methods("PUT", "OPTIONS")
	//
	//// API that needs both authentication and role-based access control (RBAC)
	//r.Handle("/delete/{id}", jwtMiddleware.Handler(AccessHandler{checkAccess, deleteHandler})).Methods("DELETE", "OPTIONS")

	return r
}