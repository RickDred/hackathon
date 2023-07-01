package fdb

import (
	"context"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go"
)

func InitDB() (*firebase.App, error) {
	opt := option.WithCredentialsFile("internal/repository/fdb/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
		return nil, err
	}

	return app, nil
}
