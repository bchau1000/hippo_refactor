package handlers

import (
	"base/server/logging"
	"base/server/models"
	"encoding/json"
	"fmt"
	"net/http"

	"firebase.google.com/go/auth"
)

type UserHandler struct {
	Client *auth.Client
}

// POST
func (handler UserHandler) LoginUser() http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var user models.User
		err := json.NewDecoder(req.Body).Decode(&user)

		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		cookie := &http.Cookie{
			Name:   "idToken",
			Value:  user.IdToken,
			MaxAge: 2592000000,
		}

		http.SetCookie(resp, cookie)

		logging.Log(fmt.Sprintf(
			"User{displayName: %s, email: %s, uid: %s} has logged in successfully",
			user.DisplayName,
			user.Email,
			user.Uid,
		))

		jsonData := []byte(fmt.Sprintf(`{"message": "%s successfully logged in!"}`, user.Email))

		resp.WriteHeader(http.StatusCreated)
		resp.Write(jsonData)
	})
}
