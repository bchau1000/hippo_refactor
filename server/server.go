package main

import (
	"base/server/config"
	"base/server/controllers"
	"base/server/logging"
	"base/server/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/**
 * TODO:
 * 1. Rename controllers to handlers
 * 2. Connect servers to Firebase auth
 * 		- Make super simple login page in front-end for testing purposes
 * 3. Performance logger (use context library?)
 * 4. Unit testing
**/
func Init() {
	startTime := time.Now()
	cfg := config.Init()
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	logging.Log(fmt.Sprintf("Starting server at %s", address))

	// Initialize controllers
	pingController := controllers.PingController{}

	// Map controllers to URLs
	router := mux.NewRouter()

	router.HandleFunc(
		"/api/version",
		middleware.Wrap(pingController.GetVersion(), middleware.ResponseHeaders()),
	).Methods("GET", "OPTIONS")

	logging.Log(fmt.Sprintf("Started server in %dms", time.Since(startTime).Milliseconds()))
	http.ListenAndServe(address, router)
}
