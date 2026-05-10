package services

import (
	"database/sql"
	"notes_api/internal/models"
	"notes_api/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *sql.DB
}

func (s *AuthService) Signup(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	err = repositories.CreateUser(s.DB, user)
	if err != nil {
		return err
	}

	return nil
}
