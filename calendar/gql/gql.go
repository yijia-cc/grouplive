package gql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yijia-cc/grouplive/calendar/dep"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/yijia-cc/grouplive/calendar/config"
)

func StartServer(cfg config.Config, db *sql.DB) {
	content, err := readStringFromFile(cfg.GraphQLSchemaPath)
	if err != nil {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(content, dep.InitGraphQLResolver(db))

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
