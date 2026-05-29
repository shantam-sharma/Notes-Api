package repositories

import (
	"database/sql"
	"errors"
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

func (r *NoteRepository) GetNotesByUserID(userID int, limit int, offset int) ([]models.Note, error) {
	query := `
		SELECT id, user_id, title, content, created_at, updated_at
		FROM notes
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
		OFFSET $3
	`

	rows, err := r.DB.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note

		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteRepository) GetNoteByID(noteID, userID int) (models.Note, error) {
	query := `
		SELECT id, user_id, title, content, created_at, updated_at
		FROM notes
		WHERE id = $1 AND user_id = $2
	`
	row := r.DB.QueryRow(query, noteID, userID)

	var note models.Note

	err := row.Scan(
		&note.ID,
		&note.UserID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func (r *NoteRepository) UpdateNote(noteID, userID int, title, content string) error {
	query := `
		UPDATE notes
		SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3 AND user_id = $4
	`
	result, err := r.DB.Exec(
		query,
		title,
		content,
		noteID,
		userID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("note not found")
	}

	return nil
}

func (r *NoteRepository) DeleteNote(noteID, userID int) error {
	query := `
		DELETE FROM notes
		WHERE id = $1 AND user_id = $2
	`

	result, err := r.DB.Exec(
		query,
		noteID,
		userID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("note not found")
	}

	return nil
}
