package fdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"hackathon/internal/models"
)

type StudentsRepo struct {
	db *firestore.CollectionRef
}

func (s StudentsRepo) GetById(ctx context.Context, id string) (models.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) GetByEmail(ctx context.Context, email string) (models.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) GetAll(ctx context.Context) ([]models.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) Create(ctx context.Context, student models.Student) (models.Student, error) {
	d := s.db.NewDoc()
	student.Id = d.ID
	_, err := d.Create(ctx, student)
	if err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (s StudentsRepo) DeleteById(ctx context.Context, id string) error {
	_, err := s.db.Doc(id).Delete(ctx)
	return err
}

func (s StudentsRepo) Update(ctx context.Context, student models.Student) (models.Student, error) {
	//TODO implement me
	panic("implement me")
}

func NewStudentsRepo(db *firestore.Client) *StudentsRepo {
	return &StudentsRepo{
		db: db.Collection(studentsCollection),
	}
}
