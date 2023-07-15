package fdb

import (
	"context"
	"hackathon/internal/models"
	"time"
)

func (p ProfessorsRepo) GetAllMessage(ctx context.Context, professorId string) ([]models.Message, error) {
	professor, err := p.GetById(ctx, professorId)
	if err != nil {
		return nil, err
	}
	return professor.Messanger, nil
}

func (p ProfessorsRepo) CreateMessage(ctx context.Context, data string, professorId string, senderId string) error {
	professor, err := p.GetById(ctx, professorId)
	if err != nil {
		return err
	}
	professor.Messanger = append(professor.Messanger, models.Message{
		SenderId: senderId,
		Data:     data,
		Time:     time.Now(),
	})
	err = p.UpdateById(ctx, professor, professorId)
	return err
}
