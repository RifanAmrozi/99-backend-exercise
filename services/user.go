package services

import (
	"99-backend-exercise/models"
	"99-backend-exercise/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}
func GetUserByID(id int) (*models.User, error) {
	return repositories.GetUserByID(id)
}
