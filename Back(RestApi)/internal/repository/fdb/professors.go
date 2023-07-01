package fdb

import "cloud.google.com/go/firestore"

type ProfessorsRepo struct {
	db *firestore.CollectionRef
}

func (p ProfessorsRepo) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) GetById() {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) GetByName() {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) Create() {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) Delete() {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) Update() {
	//TODO implement me
	panic("implement me")
}

func NewProfessorsRepo(db *firestore.Client) *ProfessorsRepo {
	return &ProfessorsRepo{
		db: db.Collection(professorsCollection),
	}
}
