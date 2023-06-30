package services

import (
	"errors"
	"fmt"
	"hackathon/api/models"
)

// RegisterUser registers a new user
func RegisterUser(user *models.Student) error {
	// Perform validation, database operations, etc.
	// Example implementation:
	if user.Email == "" || *user.Password.Plaintext == "" {
		return errors.New("email and password are required")
	}

	// Save the user to the database or perform other operations

	return nil
}

// AuthenticateUser authenticates a user with the provided email and password
func AuthenticateUser(email, password string) (string, error) {
	// Perform authentication logic, database operations, etc.
	// Example implementation:
	if email == "test@example.com" && password == "password" {
		// Generate and return a token (e.g., JWT)
		return "your-generated-token", nil
	}

	return "", fmt.Errorf("authentication failed")
}
