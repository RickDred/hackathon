package services

import (
	"context"
	"hackathon/internal/models"
	"hackathon/internal/repository"
	"hackathon/pkg/filters"
	"hackathon/pkg/validator"
	"log"
)

type ReviewsService struct {
	repo repository.Reviews
}

func NewReviewsService(repo repository.Reviews) *ReviewsService {
	return &ReviewsService{repo}
}

func (r *ReviewsService) AddNewReview(ctx context.Context, review *models.Review) (models.Review, error) {
	v := validator.New()
	if ValidateReview(v, review); !v.Valid() {
		log.Fatal("data is not valid")
	}
	return r.repo.Create(ctx, *review)

}

func (r *ReviewsService) ListReviewsByProfessor(ctx context.Context, id string, f filters.Filters) ([]models.Review, filters.Metadata, error) {
	return r.repo.GetAllByProfessor(ctx, id, f)
}
func (r *ReviewsService) UpdateReview(ctx context.Context, review *models.Review) (models.Review, error) {
	v := validator.New()
	if ValidateReview(v, review); !v.Valid() {
		log.Fatalf("data is not valid")
	}

	var err error
	if err != nil {
		log.Fatal(err)
	}
	if review.Id == "" {
		temp, _ := r.repo.GetById(ctx, review.Id)
		review.Id = temp.Id
	}
	return r.repo.UpdateById(ctx, *review, review.Id)
}
func (r *ReviewsService) DeleteReview(ctx context.Context, id string) error {
	review, err := r.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return r.repo.DeleteById(ctx, review.Id)
}
