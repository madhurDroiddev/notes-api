package repository

import "notes-api/models"

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
	GetByUser(userID int) ([]models.Note, error)
	GetByID(id, userID int) (models.Note, error)
	Update(note models.Note) error
	Delete(id, userID int) error
	Search(userID int, query string) ([]models.Note, error)
}

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
}
