package v1

import (
	"github.com/gorilla/mux"
	"hackathon/api/v1/handlers"
)

// InitRouter initializes the API router and sets up the routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Register routes
	//router.HandleFunc("/api/v1/register", handlers.RegisterUser).Methods("POST")
	//router.HandleFunc("/api/v1/login", handlers.LoginUser).Methods("POST")

	router.HandleFunc("/api/v1/healthcheck", handlers.HealthcheckHandler).Methods("GET")

	return router
}
