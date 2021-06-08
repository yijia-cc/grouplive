package server

import (
	"io/ioutil"
	"log"
	"net/http"

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

	schema := graphql.MustParseSchema(content, resolver.NewResolver(authorizer, transactionFactory, amenityDao, amenityTypeDao))
	mux := http.NewServeMux()
	relayHandler := &relay.Handler{Schema: schema}
	mux.HandleFunc("/", auth.WithMiddleware(authenticator, userProvider, relayHandler.ServeHTTP))
	return mux
}

func readStringFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
