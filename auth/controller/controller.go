package controller

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/auth/config"
	"net/http"
)

//func StartUp() *mux.Router {
//	jwtMiddleware := startUpJWT()
//	router := startUpHttpRouters(jwtMiddleware)
//	return router
//}

var SigningKey = []byte(config.AppConfig.SecretKey)

// StartupJWT adds an extra layer of security middleware between the client and the backend APIs, to protect APIs from un-authenticated users
func StartupJWT() *jwtmiddleware.JWTMiddleware {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return SigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		UserProperty:  "userToken",
	})

	return jwtMiddleware
}

func StartupHttpRouter(jwtMiddleware *jwtmiddleware.JWTMiddleware) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/post/{id}", jwtMiddleware.HandlerWithAccessControl(http.HandlerFunc(deleteHandler), checkAccess)).Methods("DELETE", "OPTIONS")
	//r.Handle("/post/{id}", jwtMiddleware.HandlerWithAccessControl(http.HandlerFunc(alertHandler), checkAccess)).Methods("DELETE", "OPTIONS")

	//r.Handle("/token", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("POST", "OPTIONS")

	r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")

	// The tokenHandler is for testing purpose only
	r.Handle("/token", jwtMiddleware.Handler(http.HandlerFunc(tokenHandler))).Methods("POST", "OPTIONS")

	return r
}
