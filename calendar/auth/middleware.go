package auth

import (
	"context"
	"net/http"
)

func WithMiddleware(authenticator Authenticator, userProvider UserProvider, handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authToken, err := TokenFromRequest(request)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		userID, err := authenticator.VerifyIdentity(authToken)
		if err != nil {
			handleFunc(writer, request)
			return
		}

		user, err := userProvider.GetUser(userID)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.Background()
		ctx = NewContextWithUser(ctx, user)
		req := request.WithContext(ctx)
		handleFunc(writer, req)
	}
}


