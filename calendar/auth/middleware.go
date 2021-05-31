package auth

import (
	"context"
	"net/http"
)

func WithMiddleware(authenticator Authenticator, handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authToken, err := TokenFromRequest(request)
		if err != nil {
			return
		}

		user := authenticator.GetUser(authToken)
		ctx := context.Background()
		ctx = NewContextWithUser(ctx, user)
		req := request.WithContext(ctx)
		handleFunc(writer, req)
	}
}
