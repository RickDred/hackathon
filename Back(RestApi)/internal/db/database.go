package db

import (
	"context"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go"
)

var (
	app *firebase.App
)

func InitDB() (*firebase.App, error) {
	opt := option.WithCredentialsFile("internal/db/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
		return nil, err
	}

	return app, nil
}
