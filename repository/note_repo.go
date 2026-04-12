package repository

import (
	"notes-api/models"

	"github.com/jmoiron/sqlx"
)

type npteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &npteRepository{db: db}
}

func (r *npteRepository) Create(note models.Note) (models.Note, error) {
	query := `INSERT INTO notes (user_id, title, content) 
              VALUES (:user_id, :title, :content) 
              RETURNING id, user_id, title, content, created_at, updated_at`
	rows, err := r.db.NamedQuery(query, note)
	if err != nil {
		return models.Note{}, err
	}
	defer rows.Close()
	rows.Next()
	rows.StructScan(&note)
	return note, nil
}

func (r *npteRepository) GetByUser(userID int) ([]models.Note, error) {
	var notes []models.Note
	err := r.db.Select(&notes, `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM notes WHERE user_id=$1 ORDER BY created_at DESC`, userID)
	return notes, err
}

func (r *npteRepository) GetByID(id, userID int) (models.Note, error) {
	var note models.Note
	err := r.db.Get(&note, `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM notes WHERE id=$1 AND user_id=$2`, id, userID)
	return note, err
}

func (r *npteRepository) Update(note models.Note) error {
	_, err := r.db.NamedExec(`UPDATE notes SET title=:title, content=:content, 
                            updated_at=NOW() WHERE id=:id AND user_id=:user_id`, note)
	return err
}

func (r *npteRepository) Delete(id, userID int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id=$1 AND user_id=$2", id, userID)
	return err
}

func (r *npteRepository) Search(userID int, query string) ([]models.Note, error) {
	var notes []models.Note
	err := r.db.Select(&notes, `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM notes 
        WHERE user_id = $1 
        AND search_vector @@ plainto_tsquery('english', $2)
        ORDER BY created_at DESC`,
		userID, query,
	)
	return notes, err
}
