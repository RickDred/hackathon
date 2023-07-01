package repository

import (
	"cloud.google.com/go/firestore"
	"hackathon/internal/repository/fdb"
)

type Professors interface {
	GetAll()
	GetById()
	GetByName()
	Create()
	Delete()
	Update()
}

type Students interface {
	GetById()
	GetByEmail()
	GetAll()
	Create()
	Delete()
	Update()
}

type Reviews interface {
	Create()
	Update()
	Delete()
	GetAllByProfessor()
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
