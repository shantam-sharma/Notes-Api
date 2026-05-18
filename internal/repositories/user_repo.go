package repositories

import (
	"database/sql"
	"notes_api/internal/models"
)

/*
|| Module Doc ||
this layer talks to postgres
this layer handle db queries
*/

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := `
		INSERT INTO users (name , email , password_hash)
		VALUES ($1 , $2 , $3)
	`
	// $ are posgres placeholders so $1 = user. name
	// _ means ignore this value for now
	_, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `
		SELECT id, name, email, password_hash
		FROM users
		WHERE email = $1
	`
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}
