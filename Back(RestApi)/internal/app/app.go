package app

import (
	v1 "hackathon/internal/delivery/http"
	"log"
	"net/http"
)

func Run() {
	//app, err := fdb.InitDB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client, err := app.Firestore(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer client.Close()
	//
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
