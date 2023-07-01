package services

import (
	"context"
	"errors"
	"hackathon/internal/models"
	"hackathon/internal/repository"
	"hackathon/pkg/validator"
)

type StudentsService struct {
	repo repository.Students
}

func NewStudentsService(repo repository.Students) *StudentsService {
	return &StudentsService{repo}
}

// RegisterStudent registers a new student
func (s *StudentsService) RegisterStudent(ctx context.Context, student *models.Student) error {
	var err error
	student.Password, err = Hash(student.Password)
	if err != nil {
		return err
	}

	v := validator.New()
	if ValidateStudent(v, student); !v.Valid() {
		return errors.New("data is not valid")
	}
	return s.repo.Create(ctx, student)
}

// AuthenticateStudent authenticates a user with the provided email and password
func (s *StudentsService) AuthenticateStudent(ctx context.Context, email, password string) (*models.Student, error) {
	student, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !CheckHash(password, student.Password) {
		return nil, errors.New("wrong password")
	}

	return &student, nil
}
