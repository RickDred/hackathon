package fdb

import (
	"cloud.google.com/go/firestore"
	"context"
	"hackathon/internal/models"
)

type ReviewsRepo struct {
	db *firestore.CollectionRef
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

func (r ReviewsRepo) Update(ctx context.Context, review models.Review) (models.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewsRepo) DeleteById(ctx context.Context, id string) error {
	_, err := r.db.Doc(id).Delete(ctx)
	return err
}

func (r ReviewsRepo) GetAllByProfessor(ctx context.Context, id string) ([]models.Review, error) {
	//TODO implement me
	panic("implement me")
}

func NewReviewsRepo(db *firestore.Client) *ReviewsRepo {
	return &ReviewsRepo{
		db: db.Collection(reviewsCollection),
	}
}
