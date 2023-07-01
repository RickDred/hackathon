package http

import (
	"github.com/gorilla/mux"
	"hackathon/internal/delivery/http/v1"
)

// InitRouter initializes the API router and sets up the routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Register routes
	//router.HandleFunc("/api/v1/register", handlers.RegisterUser).Methods("POST")
	//router.HandleFunc("/api/v1/login", handlers.LoginUser).Methods("POST")

	router.HandleFunc("/api/v1/healthcheck", v1.HealthcheckHandler).Methods("GET")

	return router
}
