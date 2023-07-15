package v1

import (
	"hackathon/internal/services"

	"github.com/gin-gonic/gin"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initStudentsRoutes(v1)
		h.initHealthCheckRoutes(v1)
		h.initProfessorsRoutes(v1)
		h.initreviewsRoutes(v1)
	}
}
