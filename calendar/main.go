package main

import (
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/yijia-cc/grouplive/calendar/gqlapi"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	content, err := readStringFromFile("gqlapi/schema.graphqls")
	if err != nil {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(content, &gqlapi.Resolver{})
	http.Handle("/", &relay.Handler{Schema: schema})

	fmt.Println("GraphQL API started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readStringFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
