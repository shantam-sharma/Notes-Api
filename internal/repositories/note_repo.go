package repositories

import (
	"database/sql"
	"notes_api/internal/models"
)

type NoteRepository struct {
	DB *sql.DB
}

func (r *NoteRepository) CreateNote(note models.Note) error {
	query := `
		INSERT INTO notes (user_id, title, content)
		VALUES ($1 , $2 , $3)
	`
	_, err := r.DB.Exec(
		query,
		note.UserID,
		note.Title,
		note.Content,
	)

	if err != nil {
		return err
	}

	return nil
}
