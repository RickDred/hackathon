package services

import "hackathon/internal/repository"

type ProfessorsService struct {
	repo repository.Professors
}

func NewProfessorsService(repo repository.Professors) *ProfessorsService {
	return &ProfessorsService{repo}
}
