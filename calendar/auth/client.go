package auth

import (
	"context"

	"github.com/yijia-cc/grouplive/calendar/auth/permission"
	"github.com/yijia-cc/grouplive/calendar/entity"
	pb "github.com/yijia-cc/grouplive/proto/golang"
	"google.golang.org/grpc"
)

var _ Authenticator = (*GroupLiveAuthClient)(nil)
var _ Authorizer = (*GroupLiveAuthClient)(nil)
var _ UserProvider = (*GroupLiveAuthClient)(nil)

type GroupLiveAuthClient struct {
	authenticationClient pb.AuthenticationServiceClient
	authorizationClient  pb.AuthorizationServiceClient
	userProviderClient   pb.UserServiceClient
}

func (g GroupLiveAuthClient) VerifyIdentity(authToken string) (string, error) {
	ctx := context.Background()
	req := &pb.VerifyIdentityRequest{AuthToken: authToken}
	res, err := g.authenticationClient.VerifyIdentity(ctx, req)
	if err != nil {
		return "", err
	}
	return res.UserId, nil
}

func (g GroupLiveAuthClient) HasPermission(user *entity.User, permission permission.Permission) bool {
	userID := "*"
	if user != nil {
		userID = string(user.ID)
	}

	ctx := context.Background()
	req := &pb.HasPermissionRequest{
		UserId:         userID,
		PermissionId:   permission.ID,
		ResourceTypeId: permission.ResourceTypeID,
		ResourceId:     permission.ResourceID,
	}
	res, err := g.authorizationClient.HasPermission(ctx, req)
	if err != nil {
		return false
	}
	return res.HasPermission
}

func (g GroupLiveAuthClient) GetUser(userID string) (entity.User, error) {
	ctx := context.Background()
	req := &pb.GetUserRequest{
		UserId: userID,
	}

	res, err := g.userProviderClient.GetUser(ctx, req)
	if err != nil {
		return entity.User{}, err
	}

	return entity.NewUserFromProto(res), err
}

func NewGroupLiveAuthClient(gRPCEndpoint string) (GroupLiveAuthClient, error) {
	conn, err := grpc.Dial(gRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		return GroupLiveAuthClient{}, err
	}
	return GroupLiveAuthClient{
		authenticationClient: pb.NewAuthenticationServiceClient(conn),
		authorizationClient:  pb.NewAuthorizationServiceClient(conn),
		userProviderClient:   pb.NewUserServiceClient(conn),
	}, nil
}
