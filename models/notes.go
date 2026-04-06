package models

type Note struct {
	ID        int    `db:"id"        json:"id"`
	UserID    int    `db:"user_id"   json:"user_id"`
	Title     string `db:"title"     json:"title"`
	Content   string `db:"content"   json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
