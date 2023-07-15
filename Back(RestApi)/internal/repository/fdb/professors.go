package fdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"hackathon/internal/models"
	"hackathon/pkg/filters"
	"hackathon/pkg/validator"
)

type ProfessorsRepo struct {
	db *firestore.CollectionRef
}

func (p ProfessorsRepo) GetAll(ctx context.Context, name string, department string, degree string, f *filters.Filters) ([]models.Professor, *filters.Metadata, error) {
	if f.PageSize == 0 {
		f.PageSize = 25
	}
	if f.Page == 0 {
		f.Page = 1
	}
	if f.SortColumn() == "" {
		f.Sort = "Id"
	}

	v := validator.New()
	filters.ValidateFilters(v, *f)

	var professors []models.Professor

	q := p.db.OrderBy(f.SortColumn(), firestore.Direction(f.SortDirection())).Limit(f.Limit()).Offset(f.Offset())

	if name != "" {
		q = q.Where("Name", "==", name)
	}
	if department != "" {
		q = q.Where("Department", "==", department)
	}
	if degree != "" {
		q = q.Where("Degree", "==", degree)
	}

	iter := q.Documents(ctx)
	defer iter.Stop()

	totalRecords := 0

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		var professor models.Professor
		if err := doc.DataTo(&professor); err != nil {
			return nil, nil, err
		}
		professors = append(professors, professor)
		totalRecords++
	}
	metadata := filters.CalculateMetadata(totalRecords, f.Page, f.PageSize)
	return professors, &metadata, nil
}

func (p ProfessorsRepo) GetById(ctx context.Context, id string) (models.Professor, error) {
	snap, err := p.db.Doc(id).Get(ctx)
	if err != nil {
		return models.Professor{}, err
	}
	data := snap.Data()
	professor := models.Professor{
		Id:         id,
		Name:       data["name"].(string),
		Email:      data["email"].(string),
		Department: data["department"].(string),
		Degree:     data["degree"].(string),
		Subjects:   data["subjects"].([]string),
	}
	return professor, nil
}

func (p ProfessorsRepo) GetByEmail(ctx context.Context, email string) (models.Professor, error) {
	iter := p.db.Where("Email", "==", email).Documents(ctx)
	defer iter.Stop()

	var professor models.Professor

	doc, err := iter.Next()
	if err != nil {
		return models.Professor{}, err
	}
	if err = doc.DataTo(&professor); err != nil {
		return models.Professor{}, err
	}

	return professor, nil
}

func (p ProfessorsRepo) Create(ctx context.Context, professor *models.Professor) error {
	d := p.db.NewDoc()
	professor.Id = d.ID
	_, err := d.Create(ctx, professor)
	return err
}

func (p ProfessorsRepo) DeleteById(ctx context.Context, id string) error {
	_, err := p.db.Doc(id).Delete(ctx)
	return err
}

func (p ProfessorsRepo) UpdateById(ctx context.Context, professor models.Professor, id string) error {
	professor.Id = id
	_, err := p.db.Doc(id).Set(ctx, professor)
	return err
}

func NewProfessorsRepo(db *firestore.Client) *ProfessorsRepo {
	return &ProfessorsRepo{
		db: db.Collection(professorsCollection),
	}
}
