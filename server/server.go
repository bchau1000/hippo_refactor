package main

import (
	"base/server/config"
	"base/server/handlers"
	"base/server/logging"
	"base/server/middleware"
	"fmt"
	"net/http"
	"net/http/cookiejar"
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
var client http.Client

func Init() {
	startTime := time.Now()
	cfg := config.Init()
	jar, err := cookiejar.New(nil)

	if err != nil {
		logging.Log(fmt.Sprintf("Got error while creating cookie jar %s", err.Error()))
	}

	client = http.Client{
		Jar: jar,
	}

	// Initialize controllers
	pingHandler := handlers.PingHandler{}
	userHandler := handlers.UserHandler{}

	// Map handlers to URLs
	router := mux.NewRouter()

	router.HandleFunc(
		"/api/version",
		middleware.Wrap(pingHandler.GetVersion(), middleware.ResponseHeaders()),
	).Methods("GET", "OPTIONS")

	router.HandleFunc(
		"/api/login",
		middleware.Wrap(userHandler.LoginUser(), middleware.ResponseHeaders()),
	).Methods("POST", "OPTIONS")

	logging.Log(fmt.Sprintf(
		"Started server at %s:%d in %dms",
		cfg.Server.Host,
		cfg.Server.Port,
		time.Since(startTime).Milliseconds()))

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router)
}
