package controllers

import (
	cfg "base/server/config"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type PingController struct {
}

// GET
func (pc PingController) GetVersion(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.WriteHeader(http.StatusOK)
	time.Sleep(2 * time.Second)
	// TO DO: Has to be a better way to get this apiVersion
	jsonData := []byte(fmt.Sprintf(
		`{"status": "OK", "version": %s}`,
		viper.GetString(cfg.ServerVersion)))

	resp.Write(jsonData)
}
