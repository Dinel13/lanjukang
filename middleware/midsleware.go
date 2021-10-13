package middleware

import (
	"fmt"
	"net/http"
)

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do something with r.Context()
		fmt.Println(r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
