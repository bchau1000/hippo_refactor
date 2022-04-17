package main

import (
	cfg "base/server/config"
	"base/server/controllers"
	"base/server/logging"
	"fmt"
	"net/http"
	"time"
)

func Init() {
	startTime := time.Now()
	conf := cfg.Init()
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	logging.Log(fmt.Sprintf("Starting server at %s", address))

	// Initialize controllers
	pingController := controllers.PingController{}

	// Map controllers to URLs
	mux := http.NewServeMux()
	mux.HandleFunc("/api/version", pingController.GetVersion)

	logging.Log(fmt.Sprintf("Server started in %dms", time.Since(startTime).Milliseconds()))
	http.ListenAndServe(address, mux)
}
