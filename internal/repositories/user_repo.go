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
func CreateUser(db *sql.DB, user models.User) error {
	query := `
		INSERT INTO users (name , email , password_hash)
		VALUES ($1 , $2 , $3)
	`
	// $ are posgres placeholders so $1 = user. name
	// _ means ignore this value for now
	_, err := db.Exec(
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
