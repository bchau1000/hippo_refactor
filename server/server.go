package main

import (
	cfg "base/server/config"
	"base/server/controllers"
	"base/server/logging"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Init() {
	startTime := time.Now()
	conf := cfg.Init()
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	logging.Log(fmt.Sprintf("Starting server at %s", address))

	// Initialize controllers
	pingController := controllers.PingController{}

	// Map controllers to URLs
	router := mux.NewRouter()
	router.HandleFunc("/api/version", pingController.GetVersion).Methods("GET", "OPTIONS")

	logging.Log(fmt.Sprintf("Started server in %dms", time.Since(startTime).Milliseconds()))
	http.ListenAndServe(address, router)
}
