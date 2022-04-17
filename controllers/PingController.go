package controllers

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type PingController struct {
}

func (pc PingController) GetVersion(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	// TO DO: Has to be a better way to get this apiVersion
	jsonData := []byte(fmt.Sprintf(`{"status": "OK", "version": %s}`, viper.GetString("apiVersion")))
	resp.Write(jsonData)
}
