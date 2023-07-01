package app

import (
	"context"
	v1 "hackathon/internal/delivery/http"
	"hackathon/internal/models"
	"hackathon/internal/repository"
	"hackathon/internal/repository/fdb"
	"log"
	"net/http"
)

func Run() {
	app, err := fdb.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	rp := repository.NewFBRepository(client)
	rp.Professors.Create(context.Background(), models.Professor{
		Name:       "NG(test)",
		Email:      "ok@gmail.com",
		Department: "calculus",
		Degree:     "2",
		Subjects: []string{
			"calc2",
			"discrete",
		},
	})

	//client.Collection("test").Add(context.Background(), map[string]interface{}{
	//	"name":    "Los Angeles",
	//	"state":   "CA",
	//	"country": "USA",
	//})
	//if err != nil {
	//	// Handle any errors in an appropriate way, such as returning them.
	//	log.Printf("An error has occurred: %s", err)
	//}

	router := v1.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
