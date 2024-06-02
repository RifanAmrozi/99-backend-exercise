package repositories

import (
	"99-backend-exercise/db"
	"99-backend-exercise/models"
	"time"
)

func CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	query := `INSERT INTO users (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`
	return db.DB.QueryRow(query, user.Name, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, created_at, updated_at FROM users WHERE id = $1`
	if err := db.DB.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}
