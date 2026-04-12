package repository

import (
	"notes-api/models"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	query := `INSERT INTO users (name, email, password) 
              VALUES (:name, :email, :password)`
	_, err := r.db.NamedExec(query, user)
	return user, err
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}
