package app

import (
	"context"
	"fmt"
	v1 "hackathon/internal/delivery/http"
	"hackathon/internal/repository"
	"hackathon/internal/repository/fdb"
	"hackathon/pkg/filters"
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

	professors, _, _ := rp.Professors.GetAll(context.Background(), filters.Filters{})

	fmt.Println(professors)

	//rp.Professors.Create(context.Background(), models.Professor{
	//	Name:       "Askar",
	//	Email:      "Askar@gmail.com",
	//	Department: "OOP",
	//	Degree:     "10",
	//	Subjects: []string{
	//		"OOP",
	//		"Advanced",
	//	},
	//})

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
