package controller

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/auth/config"
	"net/http"
)

var SigningKey []byte

func StartUp(cfg *config.Config) *mux.Router {
	SigningKey = []byte(cfg.TokenSecretKey)
	jwtMiddleware := StartupJWT()
	router := StartupHttpRouter(jwtMiddleware)
	return router
}

// StartupJWT adds an extra layer of security middleware between the client and the backend APIs, to protect APIs from disallowed users
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

// StartupHttpRouter configures the HTTP router to deliver requests to the desired APIs via URL patten matching, and at
// the same time, the JWT middleware is prepended before the router to protect 3 types of API that needs different level of protection
func StartupHttpRouter(jwtMiddleware *jwtmiddleware.JWTMiddleware) *mux.Router {
	r := mux.NewRouter()

	// 1. Free API that needs no protection
	r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")

	// 2. API that needs user authentication
	r.Handle("/token", jwtMiddleware.Handler(http.HandlerFunc(tokenHandler))).Methods("POST", "OPTIONS")

	// 3. API that needs both user authentication and user authorization (via Role-based Access Control, RBAC)
	r.Handle("/post/{id}", jwtMiddleware.Handler(AccessHandler{checkAccess, deleteHandler})).Methods("DELETE", "OPTIONS")

	return r
}