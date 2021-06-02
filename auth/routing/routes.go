package routing

import (
	"net/http"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/service"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type route struct {
	path       string
	method     string
	handleFunc http.HandlerFunc
}

func getRoutes(timer tm.Timer, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int) []route {
	authenticationService := service.NewAuthentication(timer, txFactory, userDao, jwtSigningKey, caesarCipherOffset)
	return []route{
		{
			path:       "/sign-in",
			method:     http.MethodPost,
			handleFunc: newSignInHandlerFunc(authenticationService),
		},
		{
			path:       "/sign-up",
			method:     http.MethodPost,
			handleFunc: newSignUpHandlerFunc(authenticationService),
		},
	}
}
