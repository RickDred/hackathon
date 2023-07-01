package fdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"hackathon/internal/models"
)

type ProfessorsRepo struct {
	db *firestore.CollectionRef
}

func (p ProfessorsRepo) GetAll(ctx context.Context) ([]models.Professor, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) GetById(ctx context.Context, id string) (models.Professor, error) {
	snap, err := p.db.Doc(id).Get(ctx)
	if err != nil {
		return models.Professor{}, err
	}
	return models.Professor{snap.Data()}, nil
}

func (p ProfessorsRepo) GetByName(ctx context.Context, name string) (models.Professor, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfessorsRepo) Create(ctx context.Context, professor models.Professor) (models.Professor, error) {
	d := p.db.NewDoc()
	professor.Id = d.ID
	_, err := d.Create(ctx, professor)
	if err != nil {
		return models.Professor{}, err
	}
	return professor, nil
}

func (p ProfessorsRepo) DeleteById(ctx context.Context, id string) error {
	_, err := p.db.Doc(id).Delete(ctx)
	return err
}

func (p ProfessorsRepo) UpdateById(ctx context.Context, professor models.Professor, id string) (models.Professor, error) {
	professor.Id = id
	_, err := p.db.Doc(id).Set(ctx, professor)
	if err != nil {
		return models.Professor{}, err
	}
	return professor, err
}

func NewProfessorsRepo(db *firestore.Client) *ProfessorsRepo {
	return &ProfessorsRepo{
		db: db.Collection(professorsCollection),
	}
}
