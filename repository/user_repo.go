package repository

import (
	"notes-api/models"

	"github.com/jmoiron/sqlx"
)

func CreateUser(db *sqlx.DB, user models.User) error {
	query := `INSERT INTO users (name, email, password) 
              VALUES (:name, :email, :password)`
	_, err := db.NamedExec(query, user)
	return err
}

func GetUserByEmail(db *sqlx.DB, email string) (models.User, error) {
	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}
