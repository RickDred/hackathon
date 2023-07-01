package fdb

import "cloud.google.com/go/firestore"

type ReviewsRepo struct {
	db *firestore.CollectionRef
}

func (r ReviewsRepo) Create() {
	//TODO implement me
	panic("implement me")
}

func (r ReviewsRepo) Update() {
	//TODO implement me
	panic("implement me")
}

func (r ReviewsRepo) Delete() {
	//TODO implement me
	panic("implement me")
}

func (r ReviewsRepo) GetAllByProfessor() {
	//TODO implement me
	panic("implement me")
}

func NewReviewsRepo(db *firestore.Client) *ReviewsRepo {
	return &ReviewsRepo{
		db: db.Collection(reviewsCollection),
	}
}
