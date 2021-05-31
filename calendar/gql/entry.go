package gql

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/yijia-cc/grouplive/calendar/dep"

	"github.com/yijia-cc/grouplive/calendar/config"
)

func StartServer(cfg config.Config, db *sql.DB) {
	gqlAPIServer := dep.InitGraphQLServer(cfg, db)
	fmt.Printf("GraphQL API started at port %d\n", cfg.GraphQLServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.GraphQLServerPort), gqlAPIServer))
}
