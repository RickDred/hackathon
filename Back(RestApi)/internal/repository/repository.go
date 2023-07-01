package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"hackathon/internal/models"
	"hackathon/internal/repository/fdb"
)

type Professors interface {
	GetAll(ctx context.Context) ([]models.Professor, error)
	GetById(ctx context.Context, id string) (models.Professor, error)
	GetByName(ctx context.Context, name string) (models.Professor, error)
	Create(ctx context.Context, professor models.Professor) (models.Professor, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, professor models.Professor, id string) (models.Professor, error)
}

type Students interface {
	GetById(ctx context.Context, id string) (models.Student, error)
	GetByEmail(ctx context.Context, email string) (models.Student, error)
	GetAll(ctx context.Context) ([]models.Student, error)
	Create(ctx context.Context, student models.Student) (models.Student, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, student models.Student) (models.Student, error)
}

type Reviews interface {
	Create(ctx context.Context, review models.Review) (models.Review, error)
	Update(ctx context.Context, review models.Review) (models.Review, error)
	DeleteById(ctx context.Context, id string) error
	GetAllByProfessor(ctx context.Context, id string) ([]models.Review, error)
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
