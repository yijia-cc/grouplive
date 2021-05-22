package gql

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/yijia-cc/grouplive/calendar/config"
	"github.com/yijia-cc/grouplive/calendar/gql/resolver"
	"io/ioutil"
	"log"
	"net/http"
)

func StartServer(cfg config.Config)  {
	content, err := readStringFromFile(cfg.GraphQLSchemaPath)
	if err != nil {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(content, &resolver.Resolver{})

	mux := http.NewServeMux()
	mux.Handle("/", &relay.Handler{Schema: schema})

	fmt.Printf("GraphQL API started at port %d\n", cfg.GraphQLServerPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.GraphQLServerPort), mux))
}


func readStringFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
