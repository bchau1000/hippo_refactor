package handlers

import (
	"base/server/logging"
	"context"
	"encoding/json"
	"net/http"

	"firebase.google.com/go/auth"
)

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInfo struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UserHandler struct {
	Client *auth.Client
}

// PUT
func (handler UserHandler) RegisterUser() http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var registerInfo RegisterInfo

		err := json.NewDecoder(req.Body).Decode(&registerInfo)

		if err != nil {
			logging.Log(err.Error())
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		params := (&auth.UserToCreate{}).
			DisplayName(registerInfo.DisplayName).
			Email(registerInfo.Email).
			EmailVerified(false).
			Password(registerInfo.Password).
			PhotoURL("Placeholder").
			Disabled(false)

		_, clientErr := handler.Client.CreateUser(context.Background(), params)

		if clientErr != nil {
			logging.Log(clientErr.Error())
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		jsonData := []byte(`{"data": "User successfully registered!"}`)

		resp.WriteHeader(http.StatusOK)
		resp.Write(jsonData)
	})
}

// POST
func (handler UserHandler) LoginUser() http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		var loginInfo LoginInfo

		err := json.NewDecoder(req.Body).Decode(&loginInfo)

		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		jsonData := []byte(`{"status": "OK"}`)

		resp.WriteHeader(http.StatusOK)
		resp.Write(jsonData)
	})
}
