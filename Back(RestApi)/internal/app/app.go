package app

import (
	"context"
	"hackathon/internal/delivery/http"
	"hackathon/internal/repository"
	"hackathon/internal/services"
	"log"
)

func Run() {
	// init database
	fb, err := repository.InitFBDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// take a client
	client, err := fb.Firestore(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	// get firebase repository
	rep := repository.NewFBRepository(client)

	// create new service
	s := services.NewServices(*rep)

	// create handler
	h := http.NewHandler(s.StudentService, s.ProfessorService, s.ReviewService)

	// init router
	r := h.Init()

	r.Run()

}
