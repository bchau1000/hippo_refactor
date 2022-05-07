package handlers

import (
	"net/http"
)

/**
 * Simple handler to ping the server
**/

type PingHandler struct {
}

// GET
func (ph PingHandler) GetVersion() http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(http.StatusOK)

		jsonData := []byte(`{"status": "OK", "version": "1.0"}`)

		resp.Write(jsonData)
	})
}
