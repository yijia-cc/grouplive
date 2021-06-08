package obs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func WithRequestLog(logger Logger, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err == nil {
			return
		}

		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		logger.Info(request.Context(),
			fmt.Sprintf(
				"Proto: %s, Method: %s, Path: %s, Body: %s",
				request.Proto,
				request.Method,
				request.URL,
				body))
		handlerFunc(writer, request)
	}
}
