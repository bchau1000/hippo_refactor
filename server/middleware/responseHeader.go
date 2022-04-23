package middleware

import (
	"base/server/logging"
	"net/http"
)

func ResponseHeaders() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Content-Type", "application/json")
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			logging.Log(req.RequestURI)
			next.ServeHTTP(resp, req)
		}
	}
}
