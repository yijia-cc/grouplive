package routing

import (
	"net/http"

	"github.com/yijia-cc/grouplive/auth/idgen"

	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
)

func NewServer(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int) *http.ServeMux {
	serveMux := http.NewServeMux()
	router := mux.NewRouter()
	routes := getRoutes(timer, idGenerator, txFactory, userDao, jwtSigningKey, caesarCipherOffset)
	for _, r := range routes {
		router.HandleFunc(r.path, r.handleFunc).Methods(r.method)
	}

	serveMux.HandleFunc("/", enableCORS(router.ServeHTTP))
	return serveMux
}

func enableCORS(handlerFunc http.HandlerFunc) http.HandlerFunc { // Closure
	return func(writer http.ResponseWriter, request *http.Request) { // Closure
		writer.Header().Set("Access-Control-Allow-Origin", "*")                                // Decorator
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE") // Decorator
		writer.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization") // Decorator
		if request.Method == http.MethodOptions { // Decorator
			return // Decorator
		}

		handlerFunc(writer, request) // Closure, Decorator
	}
}
