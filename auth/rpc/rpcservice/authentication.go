package rpcservice

import (
	"context"
	"errors"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/rpc/proto"
	"github.com/yijia-cc/grouplive/auth/service"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
)

var _ proto.AuthenticationServiceServer = (*Authentication)(nil)

type Authentication struct {
	proto.UnimplementedAuthenticationServiceServer
	authenticationService service.Authentication
}

func (a Authentication) VerifyIdentity(_ context.Context, request *proto.VerifyIdentityRequest) (*proto.VerifyIdentityResponse, error) {
	userID, err := a.authenticationService.VerifyIdentity(request.AuthToken)
	if err != nil {
		return nil, errors.New("identity cannot be verified")
	}
	return &proto.VerifyIdentityResponse{UserId: userID}, nil
}

func NewAuthentication(timer tm.Timer, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey string, caesarCipherOffset int) Authentication {
	return Authentication{authenticationService: service.NewAuthentication(timer, txFactory, userDao, jwtSigningKey, caesarCipherOffset)}
}
