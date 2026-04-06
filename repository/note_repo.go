package repository

import (
	"notes-api/models"

	"github.com/jmoiron/sqlx"
)

func CreateNote(db *sqlx.DB, note models.Note) (models.Note, error) {
	query := `INSERT INTO notes (user_id, title, content) 
              VALUES (:user_id, :title, :content) 
              RETURNING id, user_id, title, content, created_at, updated_at`
	rows, err := db.NamedQuery(query, note)
	if err != nil {
		return models.Note{}, err
	}
	defer rows.Close()
	rows.Next()
	rows.StructScan(&note)
	return note, nil
}

func GetNotesByUser(db *sqlx.DB, userID int) ([]models.Note, error) {
	var notes []models.Note
	err := db.Select(&notes, `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM notes WHERE user_id=$1 ORDER BY created_at DESC`, userID)
	return notes, err
}

func GetNoteByID(db *sqlx.DB, id, userID int) (models.Note, error) {
	var note models.Note
	err := db.Get(&note, `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM notes WHERE id=$1 AND user_id=$2`, id, userID)
	return note, err
}

func UpdateNote(db *sqlx.DB, note models.Note) error {
	_, err := db.NamedExec(`UPDATE notes SET title=:title, content=:content, 
                            updated_at=NOW() WHERE id=:id AND user_id=:user_id`, note)
	return err
}

func DeleteNote(db *sqlx.DB, id, userID int) error {
	_, err := db.Exec("DELETE FROM notes WHERE id=$1 AND user_id=$2", id, userID)
	return err
}

func SearchNotes(db *sqlx.DB, userID int, query string) ([]models.Note, error) {
	var notes []models.Note
	err := db.Select(&notes, `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM notes 
        WHERE user_id = $1 
        AND search_vector @@ plainto_tsquery('english', $2)
        ORDER BY created_at DESC`,
		userID, query,
	)
	return notes, err
}
