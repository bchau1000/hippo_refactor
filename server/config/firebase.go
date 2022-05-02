package config

import (
	"base/server/logging"
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase() *auth.Client {
	opt := option.WithCredentialsFile("D:\\workspace\\hippo-3b249-firebase-adminsdk-jz8r8-cbeabdae53.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		logging.Log(fmt.Sprintf("Error initializing Firebase SDK: %v", err))
		panic(fmt.Sprintf("Error initializing Firebase SDK: %v", err))
	}

	auth, err := app.Auth(context.Background())

	if err != nil {
		logging.Log(fmt.Sprintf("Error initializing Firebase Auth: %v", err))
		panic(fmt.Sprintf("Error initializing Firebase Auth: %v", err))
	}

	return auth
}
