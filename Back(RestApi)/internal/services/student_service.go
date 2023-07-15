package services

import (
	"context"
	"errors"
	"hackathon/internal/models"
	"hackathon/internal/repository"
	"hackathon/pkg/filters"
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
	v := validator.New()
	if ValidateEmail(v, email); !v.Valid() {
		return nil, errors.New("email is not valid")
	}

	student, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !CheckHash(password, student.Password) {
		return nil, errors.New("wrong password")
	}

	return student, nil
}

func (s *StudentsService) DeleteStudent(ctx context.Context, email string) error {
	v := validator.New()
	if ValidateEmail(v, email); !v.Valid() {
		return errors.New("email is not valid")
	}

	student, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	return s.repo.DeleteById(ctx, student.Id)
}

// UpdateStudent update the user by email, so we don't update the email
func (s *StudentsService) UpdateStudent(ctx context.Context, student *models.Student) error {
	v := validator.New()
	if ValidateStudent(v, student); !v.Valid() {
		return errors.New("data is not valid")
	}

	var err error
	student.Password, err = Hash(student.Password)
	if err != nil {
		return err
	}
	if student.Id == "" {
		temp, _ := s.repo.GetByEmail(ctx, student.Email)
		student.Id = temp.Id
	}
	return s.repo.UpdateById(ctx, student, student.Id)
}

func (s *StudentsService) ListStudents(ctx context.Context, name string, email string, f *filters.Filters) ([]models.Student, *filters.Metadata, error) {
	return s.repo.GetAll(ctx, name, email, f)
}
