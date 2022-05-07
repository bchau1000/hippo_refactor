package main

import (
	"base/server/config"
	"base/server/handlers"
	"base/server/logging"
	"base/server/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	// Initialize controllers
	pingHandler := handlers.PingHandler{}

	// Map handlers to URLs
	router := mux.NewRouter()

	router.HandleFunc(
		"/api/version",
		middleware.Wrap(pingHandler.GetVersion(), middleware.ResponseHeaders()),
	).Methods("GET", "OPTIONS")

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{cfg.Server.Host, "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	logging.Log(fmt.Sprintf(
		"Started server at %s:%d in %dms",
		cfg.Server.Host,
		cfg.Server.Port,
		time.Since(startTime).Milliseconds()))

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), c.Handler(router))
}
