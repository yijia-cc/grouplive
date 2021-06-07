package rpcservice

import (
	"context"

	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/entity"
	"github.com/yijia-cc/grouplive/auth/service"
	"github.com/yijia-cc/grouplive/auth/tx"
	proto "github.com/yijia-cc/grouplive/proto/golang"
)

var _ proto.AuthorizationServiceServer = (*Authorization)(nil)

type Authorization struct {
	proto.UnimplementedAuthorizationServiceServer
	authorizationService service.Authorization
}

func (a Authorization) HasPermission(_ context.Context, request *proto.HasPermissionRequest) (*proto.HasPermissionResponse, error) {
	permission := entity.Permission{ID: (entity.ID)(request.PermissionId)}
	user := entity.User{ID: (entity.ID)(request.UserId)}
	resourceType := entity.ResourceType{ID: (entity.ID)(request.ResourceTypeId)}
	resource := entity.Resource{
		ID:   (entity.ID)(request.ResourceId),
		Type: resourceType,
	}
	hasPermission, err := a.authorizationService.HasPermission(permission, user, resource)
	if err != nil {
		return nil, err
	}
	return &proto.HasPermissionResponse{HasPermission: hasPermission}, nil
}

func NewAuthorization(txFactory tx.TransactionFactory, permissionBindingDao dao.PermissionBinding) Authorization {
	return Authorization{authorizationService: service.NewAuthorization(txFactory, permissionBindingDao)}
}
