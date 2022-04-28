package handlers

import (
	cfg "base/server/config"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
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

		// Only here to test loading/caching in frontend
		time.Sleep(2 * time.Second)

		jsonData := []byte(fmt.Sprintf(
			`{"status": "OK", "version": %s}`,
			viper.GetString(cfg.ServerVersion)))

		resp.Write(jsonData)
	})
}
