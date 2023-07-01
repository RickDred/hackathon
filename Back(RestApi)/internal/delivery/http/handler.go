package http

import (
	"github.com/gin-gonic/gin"
	v1 "hackathon/internal/delivery/http/v1"
	"hackathon/internal/services"
)

type Handler struct {
	studentService   services.StudentService
	professorService services.ProfessorService
	reviewService    services.ReviewService
}

func NewHandler(studentService services.StudentService, professorService services.ProfessorService, reviewService services.ReviewService) *Handler {
	return &Handler{
		studentService:   studentService,
		professorService: professorService,
		reviewService:    reviewService,
	}
}

// Init initializes the API router and sets up the routes
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.InitApi(router)

	return router

	// Router on mux (old version), now we use gin router

	//router := mux.NewRouter()
	//
	//// Auth routes
	//router.HandleFunc("/api/v1/register", v1.RegisterStudent).Methods("POST")
	//router.HandleFunc("/api/v1/login", v1.LoginStudent).Methods("POST")
	//
	//// health-checker. My pride, it always works
	//router.HandleFunc("/api/v1/healthcheck", v1.HealthcheckHandler).Methods("GET")
	//
	//return router
}

func (h *Handler) InitApi(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.studentService, h.professorService, h.reviewService)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
