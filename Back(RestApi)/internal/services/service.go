package services

import (
	"context"
	"hackathon/internal/models"
	"hackathon/internal/repository"
)

type StudentService interface {
	RegisterStudent(ctx context.Context, student *models.Student) error
	AuthenticateStudent(ctx context.Context, email, password string) (*models.Student, error)
}

type ProfessorService interface {
}

type ReviewService interface {
}

type Services struct {
	StudentService   StudentService
	ProfessorService ProfessorService
	ReviewService    ReviewService
}

func NewServices(r repository.Repositories) *Services {
	return &Services{
		StudentService:   NewStudentsService(r.Students),
		ProfessorService: NewProfessorsService(r.Professors),
		ReviewService:    NewReviewsService(r.Reviews),
	}
}
