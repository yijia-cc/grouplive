package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/yijia-cc/grouplive/calendar/entity"
)

type key int

const userKey key = 1

func TokenFromRequest(req *http.Request) (string, error) {
	// Expected authorization header
	// Authorization: Bearer <JWTToken>
	authHeader := req.Header.Get("Authorization")
	if len(authHeader) == 0 {
		return "", nil
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("authorization header is mal-formatted")
	}

	if parts[0] != "Bearer" {
		return "", errors.New("unsupported auth credential")
	}

	return parts[1], nil
}

func NewContextWithUser(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, userKey, &user)
}

func UserFromContext(ctx context.Context) (*entity.User, error) {
	ctxUser := ctx.Value(userKey)
	user, ok := ctxUser.(*entity.User)
	if !ok {
		return nil, errors.New("cannot cast value to user")
	}
	return user, nil
}
