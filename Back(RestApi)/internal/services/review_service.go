package services

import "hackathon/internal/repository"

type ReviewsService struct {
	repo repository.Reviews
}

func NewReviewsService(repo repository.Reviews) *ReviewsService {
	return &ReviewsService{repo}
}
