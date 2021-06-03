package routingentry

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/yijia-cc/grouplive/auth/config"
	"github.com/yijia-cc/grouplive/auth/dep"
)

func StartServer(cfg config.Config, sqlDB *sql.DB) {
	routingServer := dep.InitRoutingServer(
		dep.JWTSigningKey(cfg.JWTSigningKey),
		dep.CaesarCipherOffset(cfg.CaesarCipherOffset),
		sqlDB,
	)
	fmt.Println("GraphQL API started at port 8080")
	if err := http.ListenAndServe(":8080", routingServer); err != nil {
		panic(err)
	}
}
