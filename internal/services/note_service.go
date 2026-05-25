package services

import (
	"errors"
	"notes_api/internal/models"
	"notes_api/internal/repositories"
)

type NoteService struct {
	Repo *repositories.NoteRepository
}

func (s *NoteService) CreateNote(userID int, title, content string) error {
	if title == "" {
		return errors.New("Title is required")
	}

	note := models.Note{
		UserID:  userID,
		Title:   title,
		Content: content,
	}
	return s.Repo.CreateNote(note)
}
