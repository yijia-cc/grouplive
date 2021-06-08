package gqlentry

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/yijia-cc/grouplive/calendar/obs"

	"github.com/yijia-cc/grouplive/calendar/config"
	"github.com/yijia-cc/grouplive/calendar/dep"
)

func StartServer(cfg config.Config, logger obs.Logger, db *sql.DB) {
	gqlAPIServer, err := dep.InitGraphQLServer(cfg, logger.NextLayer().NextLayer(), db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GraphQL API started at port %d\n", cfg.GraphQLServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.GraphQLServerPort), gqlAPIServer))
}
