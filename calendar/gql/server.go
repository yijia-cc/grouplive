package gql

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yijia-cc/grouplive/calendar/obs"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/config"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/gql/resolver"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

func NewServer(
	cfg config.Config,
	logger obs.Logger,
	authenticator auth.Authenticator,
	authorizer auth.Authorizer,
	userProvider auth.UserProvider,
	transactionFactory tx.TransactionFactory,
	amenityDao dao.Amenity,
	amenityTypeDao dao.AmenityType,
) *http.ServeMux {
	content, err := readStringFromFile(cfg.GraphQLSchemaPath)
	if err != nil {
		log.Fatal(err)
	}
	res := resolver.NewResolver(logger.NextLayer(), authorizer, transactionFactory, amenityDao, amenityTypeDao)
	schema := graphql.MustParseSchema(content, res)
	mux := http.NewServeMux()

	relayHandler := &relay.Handler{Schema: schema}
	handlerFunc := auth.WithMiddleware(authenticator, userProvider, relayHandler.ServeHTTP)
	handlerFunc = obs.WithRequestLog(logger.NextLayer(), handlerFunc)

	mux.HandleFunc("/", handlerFunc)
	return mux
}

func readStringFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
