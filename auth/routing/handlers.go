package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yijia-cc/grouplive/auth/service"
)

func newSignUpHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func newSignInHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonReqBody := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		err = json.Unmarshal(buf, &jsonReqBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		authToken, err := authenticationService.SignIn(jsonReqBody.Username, jsonReqBody.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		jsonResBody := struct {
			AuthToken string `json:"auth_token"`
		}{
			AuthToken: authToken,
		}
		buf, err = json.Marshal(jsonResBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
}
