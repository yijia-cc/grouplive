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
		UserProperty: "userToken",
	})

	return jwtMiddleware
}


func StartupHttpRouter(jwtMiddleware *jwtmiddleware.JWTMiddleware) *mux.Router {
	r := mux.NewRouter()

	// 只有需要 authentication 的 API 才需要包裹一层  JWT middleware，拦截非授权用户的访问。所以 /upload, /search, /post/{id}
	// 都需要受到保护，而/signup, /signin 不需要保护
	//r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS")
	//r.Handle("/post/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE", "OPTIONS")

	r.Handle("/post/{id}", jwtMiddleware.HandlerWithAccessControl(http.HandlerFunc(deleteHandler), checkAccess)).Methods("DELETE", "OPTIONS")

	r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")


	// The tokenHandler is for testing purpose only
	r.Handle("/token", jwtMiddleware.Handler(http.HandlerFunc(tokenHandler))).Methods("POST", "OPTIONS")

	return r
}