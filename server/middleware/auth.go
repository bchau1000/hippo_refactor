package middleware

import (
	"base/server/logging"
	"base/server/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func authenticate() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			var user models.User
			err := json.NewDecoder(req.Body).Decode(&user)

			if err != nil {
				logging.Log(fmt.Sprintf("Error while authenticating user: %s", err.Error()))
				http.Error(resp, err.Error(), http.StatusBadRequest)
				return
			}

			logging.Log(fmt.Sprintf(
				"Authenticated user: {displayName: %s, email: %s, uid: %s, idToken: %s}",
				user.DisplayName,
				user.Email,
				user.Uid,
				user.IdToken,
			))

			next.ServeHTTP(resp, req)
		}
	}
}
