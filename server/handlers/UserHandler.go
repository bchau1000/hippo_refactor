package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * Simple handler to ping the server
**/

type LoginInfo struct {
	Email    string
	Password string
}

type UserHandler struct {
}

// POST
func (ph UserHandler) RegisterUser() http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(http.StatusOK)

		jsonData := []byte(fmt.Sprintf(
			`{"status": "OK", "version": %s}`,
			""))

		resp.Write(jsonData)
	})
}

func (ph UserHandler) LoginUser() http.HandlerFunc {
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
