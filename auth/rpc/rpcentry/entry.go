package rpcentry

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/yijia-cc/grouplive/auth/config"
	"github.com/yijia-cc/grouplive/auth/dep"
)

func StartServer(cfg config.Config, sqlDB *sql.DB) {
	server := dep.InitGRPCServer(dep.JWTSigningKey(cfg.JWTSigningKey),
		dep.CaesarCipherOffset(cfg.CaesarCipherOffset),
		sqlDB,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAPIPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("GRPC server started at port %d\n", cfg.GRPCAPIPort)
	panic(server.Serve(lis))
}
