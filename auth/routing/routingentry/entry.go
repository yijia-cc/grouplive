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
	fmt.Printf("Web server started at port %d\n", cfg.WebAPIPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.WebAPIPort), routingServer); err != nil {
		panic(err)
	}
}
