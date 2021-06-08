package rpcservice

import (
	"context"
	"errors"

	pb "github.com/yijia-cc/grouplive/proto/golang"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/idgen"
	"github.com/yijia-cc/grouplive/auth/service"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
)

var _ pb.AuthenticationServiceServer = (*Authentication)(nil)

type Authentication struct {
	pb.UnimplementedAuthenticationServiceServer
	authenticationService service.Authentication
}

func (a Authentication) VerifyIdentity(_ context.Context, request *pb.VerifyIdentityRequest) (*pb.VerifyIdentityResponse, error) {
	userID, err := a.authenticationService.VerifyIdentity(request.AuthToken)
	if err != nil {
		return nil, errors.New("identity cannot be verified")
	}
	return &pb.VerifyIdentityResponse{UserId: userID}, nil
}

func NewAuthentication(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int) Authentication {
	return Authentication{authenticationService: service.NewAuthentication(timer, idGenerator, txFactory, userDao, jwtSigningKey, caesarCipherOffset)}
}
