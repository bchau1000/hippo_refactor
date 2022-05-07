package handlers

import (
	"encoding/json"
	"net/http"

	"firebase.google.com/go/auth"
)

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	Client *auth.Client
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
