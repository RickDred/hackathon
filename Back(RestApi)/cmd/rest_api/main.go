package main

import (
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	v1 "hackathon/api/v1"
	"hackathon/internal/db"
	"log"
	"net/http"
)

func main() {
	// Initialize the API router
	router := v1.InitRouter()

	app, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	res := client.Collection("test").Documents(context.Background())

	for {
		doc, err := res.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
