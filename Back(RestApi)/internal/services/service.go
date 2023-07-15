package services

import (
	"context"
	"hackathon/internal/models"
	"hackathon/internal/repository"
	"hackathon/pkg/filters"
)

type StudentService interface {
	RegisterStudent(ctx context.Context, student *models.Student) error
	AuthenticateStudent(ctx context.Context, email, password string) (*models.Student, error)
	DeleteStudent(ctx context.Context, email string) error
	UpdateStudent(ctx context.Context, student *models.Student) error
	ListStudents(ctx context.Context, name string, email string, f *filters.Filters) ([]models.Student, *filters.Metadata, error)
}

type ProfessorService interface {
	AddProfessor(ctx context.Context, professor *models.Professor) error
	UpdateProfessor(ctx context.Context, professor *models.Professor) error
	GetProfessorById(ctx context.Context, id string) (*models.Professor, error)
	GetProfessorByEmail(ctx context.Context, email string) (*models.Professor, error)
	DeleteProfessor(ctx context.Context, id string) error
	ListProfessors(ctx context.Context, name string, Department string, Degree string, f *filters.Filters) ([]models.Professor, *filters.Metadata, error)
	AddMessage(ctx context.Context, data string, professorId string, senderId string) error
	ListMessages(ctx context.Context, senderId string) ([]models.Message, error)
}

type ReviewService interface {
	AddNewReview(ctx context.Context, review *models.Review) (models.Review, error)
	ListReviewsByProfessor(ctx context.Context, id string, f filters.Filters) ([]models.Review, filters.Metadata, error)
	UpdateReview(ctx context.Context, review *models.Review) (models.Review, error)
	DeleteReview(ctx context.Context, id string) error
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
