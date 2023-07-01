package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"hackathon/internal/models"
	"hackathon/internal/repository/fdb"
	"hackathon/pkg/filters"
	"log"
)

type Professors interface {
	GetAll(ctx context.Context, filters filters.Filters) ([]models.Professor, filters.Metadata, error)
	GetById(ctx context.Context, id string) (models.Professor, error)
	GetByName(ctx context.Context, name string) (models.Professor, error)
	Create(ctx context.Context, professor models.Professor) (models.Professor, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, professor models.Professor, id string) (models.Professor, error)
}

type Students interface {
	GetById(ctx context.Context, id string) (models.Student, error)
	GetByEmail(ctx context.Context, email string) (models.Student, error)
	GetAll(ctx context.Context, filters filters.Filters) ([]models.Student, filters.Metadata, error)
	Create(ctx context.Context, student *models.Student) error
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, student models.Student, id string) (models.Student, error)
}

type Reviews interface {
	Create(ctx context.Context, review models.Review) (models.Review, error)
	UpdateById(ctx context.Context, review models.Review, id string) (models.Review, error)
	GetById(ctx context.Context, id string) (models.Review, error)
	DeleteById(ctx context.Context, id string) error
	GetAllByProfessor(ctx context.Context, id string, filters filters.Filters) ([]models.Review, filters.Metadata, error)
}

type Repositories struct {
	Reviews    Reviews
	Professors Professors
	Students   Students
}

func NewFBRepository(db *firestore.Client) *Repositories {
	return &Repositories{
		Students:   fdb.NewStudentsRepo(db),
		Professors: fdb.NewProfessorsRepo(db),
		Reviews:    fdb.NewReviewsRepo(db),
	}
}

func InitFBDB() (*firebase.App, error) {
	opt := option.WithCredentialsFile("internal/repository/fdb/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
		return nil, err
	}

	return app, nil
}
