package rpc

import (
	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/idgen"
	"github.com/yijia-cc/grouplive/auth/rpc/proto"
	"github.com/yijia-cc/grouplive/auth/rpc/rpcservice"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
	"google.golang.org/grpc"
)

func NewServer(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int, permissionBinding dao.PermissionBinding) *grpc.Server {
	server := grpc.NewServer()
	proto.RegisterAuthenticationServiceServer(server, rpcservice.NewAuthentication(timer, idGenerator, txFactory, userDao, jwtSigningKey, caesarCipherOffset))
	proto.RegisterAuthorizationServiceServer(server, rpcservice.NewAuthorization(txFactory, permissionBinding))
	proto.RegisterUserServiceServer(server, rpcservice.NewUser(txFactory, userDao))
	return server
}
