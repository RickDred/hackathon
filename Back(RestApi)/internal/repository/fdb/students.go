package fdb

import (
	"cloud.google.com/go/firestore"
)

type StudentsRepo struct {
	db *firestore.CollectionRef
}

func (s StudentsRepo) GetById() {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) GetByEmail() {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) Create() {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) Delete() {
	//TODO implement me
	panic("implement me")
}

func (s StudentsRepo) Update() {
	//TODO implement me
	panic("implement me")
}

func NewStudentsRepo(db *firestore.Client) *StudentsRepo {
	return &StudentsRepo{
		db: db.Collection(studentsCollection),
	}
}
