package fdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"hackathon/internal/models"
	"hackathon/pkg/filters"
	"hackathon/pkg/validator"
)

type StudentsRepo struct {
	db *firestore.CollectionRef
}

func (s StudentsRepo) GetById(ctx context.Context, id string) (models.Student, error) {
	snap, err := s.db.Doc(id).Get(ctx)
	if err != nil {
		return models.Student{}, err
	}
	data := snap.Data()
	student := models.Student{
		Id:       id,
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}
	return student, nil
}

func (s StudentsRepo) GetByEmail(ctx context.Context, email string) (models.Student, error) {
	iter := s.db.Where("Email", "==", email).Documents(ctx)
	defer iter.Stop()

	var student models.Student

	doc, err := iter.Next()
	if err != nil {
		return models.Student{}, err
	}
	if err = doc.DataTo(&student); err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func (s StudentsRepo) GetAll(ctx context.Context, f filters.Filters) ([]models.Student, filters.Metadata, error) {
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
	filters.ValidateFilters(v, f)

	var students []models.Student

	iter := s.db.OrderBy(f.SortColumn(), firestore.Direction(f.SortDirection())).Limit(f.Limit()).Offset(f.Offset()).Documents(ctx)
	defer iter.Stop()

	totalRecords := 0

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, filters.Metadata{}, err
		}
		var student models.Student
		if err := doc.DataTo(&student); err != nil {
			return nil, filters.Metadata{}, err
		}
		students = append(students, student)
		totalRecords++
	}
	metadata := filters.CalculateMetadata(totalRecords, f.Page, f.PageSize)
	return students, metadata, nil
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

func (s StudentsRepo) UpdateById(ctx context.Context, student models.Student, id string) (models.Student, error) {
	student.Id = id
	_, err := s.db.Doc(id).Set(ctx, student)
	if err != nil {
		return models.Student{}, err
	}
	return student, err
}

func NewStudentsRepo(db *firestore.Client) *StudentsRepo {
	return &StudentsRepo{
		db: db.Collection(studentsCollection),
	}
}
