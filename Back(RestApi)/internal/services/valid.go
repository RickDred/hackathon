package services

import (
	"hackathon/internal/models"
	"hackathon/pkg/validator"
)

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 4, "password", "must be at least 4 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
func ValidateStudent(v *validator.Validator, student *models.Student) {
	v.Check(student.Name != "", "name", "must be provided")
	v.Check(len(student.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, student.Email)
	ValidatePasswordPlaintext(v, student.Password)
}
func ValidateReview(v *validator.Validator, review *models.Review) {
	v.Check(review.Feedback != "", "feedback", "must be provided")
	v.Check(review.Grading != nil, "grading", "must be provided")
	v.Check(review.Materials != nil, "materials", "must be provided")
	v.Check(review.Overall != nil, "overall", "must be provided")
}
