package obs

import (
	"bytes"
	"fmt"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"io/ioutil"
	"net/http"
)

func WithRequestLog(logger Logger, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		ctx := request.Context()

		user, err := auth.UserFromContext(ctx)
		if err != nil {
			user = nil
		}

		logger.Info(request.Context(),
			fmt.Sprintf(
				"Proto: %s, Method: %s, Path: %s, Body: %s, User: %s",
				request.Proto,
				request.Method,
				request.URL,
				body,
				user))
		handlerFunc(writer, request)
	}
}
