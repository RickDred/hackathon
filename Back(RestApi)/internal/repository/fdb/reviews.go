package fdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"hackathon/internal/models"
	"hackathon/pkg/filters"
	"hackathon/pkg/validator"
)

type ReviewsRepo struct {
	db *firestore.CollectionRef
}

func (r ReviewsRepo) GetById(ctx context.Context, id string) (models.Review, error) {
	snap, err := r.db.Doc(id).Get(ctx)
	if err != nil {
		return models.Review{}, err
	}
	data := snap.Data()
	review := models.Review{
		Id:             id,
		StudentId:      data["student_id"].(int64),
		ProfessorId:    data["professor_id"].(int64),
		Communication:  data["communication"].([]int64),
		TimeManagement: data["time_management"].([]int64),
		Grading:        data["grading"].([]int64),
		Materials:      data["materials"].([]int64),
		Overall:        data["overall"].([]int64),
		Feedback:       data["feedback"].(string),
	}
	return review, nil
}

func (r ReviewsRepo) Create(ctx context.Context, review models.Review) (models.Review, error) {
	d := r.db.NewDoc()
	review.Id = d.ID
	_, err := d.Create(ctx, review)
	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (r ReviewsRepo) UpdateById(ctx context.Context, review models.Review, id string) (models.Review, error) {
	review.Id = id
	_, err := r.db.Doc(id).Set(ctx, review)
	if err != nil {
		return models.Review{}, err
	}
	return review, err
}

func (r ReviewsRepo) DeleteById(ctx context.Context, id string) error {
	_, err := r.db.Doc(id).Delete(ctx)
	return err
}

func (r ReviewsRepo) GetAllByProfessor(ctx context.Context, id string, f filters.Filters) ([]models.Review, filters.Metadata, error) {
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

	var reviews []models.Review

	iter := r.db.Where("ProfessorId", "==", id).OrderBy(f.SortColumn(), firestore.Direction(f.SortDirection())).Limit(f.Limit()).Offset(f.Offset()).Documents(ctx)
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
		var review models.Review
		if err := doc.DataTo(&review); err != nil {
			return nil, filters.Metadata{}, err
		}
		reviews = append(reviews, review)
		totalRecords++
	}
	metadata := filters.CalculateMetadata(totalRecords, f.Page, f.PageSize)
	return reviews, metadata, nil
}

func NewReviewsRepo(db *firestore.Client) *ReviewsRepo {
	return &ReviewsRepo{
		db: db.Collection(reviewsCollection),
	}
}
