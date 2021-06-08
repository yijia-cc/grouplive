package rpc

import (
	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/idgen"
	"github.com/yijia-cc/grouplive/auth/rpc/rpcservice"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
	"github.com/yijia-cc/grouplive/proto/golang"
	"google.golang.org/grpc"
)

func NewServer(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int, permissionBinding dao.PermissionBinding) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(server, rpcservice.NewAuthentication(timer, idGenerator, txFactory, userDao, jwtSigningKey, caesarCipherOffset))
	pb.RegisterAuthorizationServiceServer(server, rpcservice.NewAuthorization(txFactory, permissionBinding))
	pb.RegisterUserServiceServer(server, rpcservice.NewUser(txFactory, userDao))
	return server
}
