package middleware

import (
	"net/http"
)

// TODO
func authenticate() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {

			next.ServeHTTP(resp, req)
		}
	}
}
