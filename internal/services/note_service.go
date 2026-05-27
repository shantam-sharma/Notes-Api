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

func (s *NoteService) GetNotesByUserID(userID int) ([]models.Note, error) {
	return s.Repo.GetNotesByUserID(userID)
}

func (s *NoteService) GetNoteByID(noteID, userID int) (models.Note, error) {
	return s.Repo.GetNoteByID(noteID, userID)
}

func (s *NoteService) UpdateNote(noteID, userID int, title, content string) error {
	return s.Repo.UpdateNote(noteID, userID, title, content)
}

func (s *NoteService) DeleteNote(noteID, UserID int) error {
	return s.Repo.DeleteNote(noteID, UserID)
}
