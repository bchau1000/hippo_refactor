package middleware

import (
	"base/server/logging"
	"fmt"
	"net/http"
)

func ResponseHeaders() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			resp.Header().Set("Access-Control-Allow-Credentials", "true")

			if req.Method == "OPTIONS" {
				resp.WriteHeader(http.StatusOK)
				resp.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, "options OK!")))
				return
			}

			logging.Log(req.RequestURI)
			next.ServeHTTP(resp, req)
		}
	}
}
