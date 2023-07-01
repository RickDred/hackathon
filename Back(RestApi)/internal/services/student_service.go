package services

import (
	"errors"
	"fmt"
	"hackathon/internal/models"
)

// RegisterStudent registers a new student
func RegisterStudent(user *models.Student) error {
	if user.Email == "" || *user.Password.Plaintext == "" {
		return errors.New("email and password are required")
	}
	return nil
}

// AuthenticateStudent authenticates a user with the provided email and password
func AuthenticateStudent(email, password string) (string, error) {

	return "", fmt.Errorf("authentication failed")
}
