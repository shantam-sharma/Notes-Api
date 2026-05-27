package services

import (
	"errors"
	"notes_api/internal/models"
	"notes_api/internal/repositories"
	"notes_api/internal/utils"

	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
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
	//! important concept to convert generic error into: PostgreSQL-specific error
	err = s.UserRepo.CreateUser(user)
	if err != nil {
		//!checks unique constraint violation
		pgErr, ok := err.(*pgconn.PgError)

		if ok && pgErr.Code == "23505" {
			return errors.New("user already exists")
		}
		return err
	}

	return nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
