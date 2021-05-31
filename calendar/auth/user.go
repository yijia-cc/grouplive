package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/yijia-cc/grouplive/calendar/entity"
)

type key int

const userKey key = 1

func TokenFromRequest(req *http.Request) (string, error) {
	panic("implement me")
}

func NewContextWithUser(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func UserFromContext(ctx context.Context) (*entity.User, error) {
	user, ok := ctx.Value(userKey).(*entity.User)
	if !ok {
		return nil, errors.New("cannot cast value to user")
	}
	return user, nil
}
