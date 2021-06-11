package auth

import (
	"fmt"
	"net/http"
)

func WithMiddleware(authenticator Authenticator, userProvider UserProvider, handleFunc http.HandlerFunc) http.HandlerFunc {
	//fmt.Println("Auth Middleware is running.....")

	return func(writer http.ResponseWriter, request *http.Request) {
		authToken, err := TokenFromRequest(request)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		//fmt.Println("authToken = ", authToken)

		userID, err := authenticator.VerifyIdentity(authToken)
		if err != nil {
			fmt.Println(err)
			//handleFunc(writer, request)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		//fmt.Println("userID = ", userID)

		user, err := userProvider.GetUser(userID)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}


		//fmt.Println("user = ", user)

		//ctx := context.Background()
		ctx := request.Context()

		ctx = NewContextWithUser(ctx, user)
		req := request.WithContext(ctx)
		handleFunc(writer, req)
	}
}
