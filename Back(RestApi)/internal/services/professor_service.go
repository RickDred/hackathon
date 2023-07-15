package services

import (
	"context"
	"hackathon/internal/models"
	"hackathon/internal/repository"
	"hackathon/pkg/filters"
)

type ProfessorsService struct {
	repo repository.Professors
}

func NewProfessorsService(repo repository.Professors) *ProfessorsService {
	return &ProfessorsService{repo}
}

func (p ProfessorsService) AddProfessor(ctx context.Context, professor *models.Professor) error {
	return p.repo.Create(ctx, professor)
}

func (p ProfessorsService) GetProfessorById(ctx context.Context, id string) (*models.Professor, error) {
	professor, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &professor, nil
}

func (p ProfessorsService) UpdateProfessor(ctx context.Context, professor *models.Professor) error {
	if professor.Id == "" {
		temp, _ := p.repo.GetByEmail(ctx, professor.Email)
		professor.Id = temp.Id
	}
	return p.repo.UpdateById(ctx, *professor, professor.Id)
}

func (p ProfessorsService) GetProfessorByEmail(ctx context.Context, email string) (*models.Professor, error) {
	professor, err := p.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &professor, nil
}

func (p ProfessorsService) DeleteProfessor(ctx context.Context, id string) error {
	return p.repo.DeleteById(ctx, id)
}

func (p ProfessorsService) ListProfessors(ctx context.Context, name string, department string, degree string, f *filters.Filters) ([]models.Professor, *filters.Metadata, error) {
	return p.repo.GetAll(ctx, name, department, degree, f)
}

func (p ProfessorsService) AddMessage(ctx context.Context, data string, professorId string, senderId string) error {
	return p.repo.CreateMessage(ctx, data, professorId, senderId)
}

func (p ProfessorsService) ListMessages(ctx context.Context, professorId string) ([]models.Message, error) {
	return p.repo.GetAllMessage(ctx, professorId)
}
