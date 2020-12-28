package behavior

import (
	"fmt"
	"net/http"
	"testing"
)

func ParamHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("params handler")
		h.ServeHTTP(w, r)
	})
}

func SignHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("sign handler")
		h.ServeHTTP(w, r)
	})
}

func TestMiddlewareHandlers_Run(t *testing.T) {
	handlers := NewMiddlewareHandlers()
	handlers.AddHandler(ParamHandler)
	handlers.AddHandler(SignHandler)

	handlers.Run(":8888")
}
