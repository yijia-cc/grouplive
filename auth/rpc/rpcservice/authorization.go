package rpcservice

import (
	"context"

	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/rpc/proto"
	"github.com/yijia-cc/grouplive/auth/service"
)

var _ proto.AuthorizationServiceServer = (*Authorization)(nil)

type Authorization struct {
	proto.UnimplementedAuthorizationServiceServer
	authorizationService service.Authorization
}

func (a Authorization) HasPermission(_ context.Context, request *proto.HasPermissionRequest) (*proto.HasPermissionResponse, error) {
	hasPermission := a.authorizationService.HasPermission(request.PermissionId, (*entity.ID)(request.UserId), (*entity.ID)(request.ResourceId))
	return &proto.HasPermissionResponse{HasPermission: hasPermission}, nil
}

func NewAuthorization() Authorization {
	return Authorization{authorizationService: service.NewAuthorization()}
}
