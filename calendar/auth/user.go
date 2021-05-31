package auth

import (
	"context"
	"errors"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"net/http"
)

type key int

const userKey key = 1

func TokenFromRequest(req *http.Request) (string, error) {
	panic("implement me")
}

func NewContext(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func UserFromContext(ctx context.Context) (*entity.User, error) {
	user, ok := ctx.Value(userKey).(*entity.User)
	if !ok {
		return nil, errors.New("cannot cast value to user")
	}
	return user, nil
}