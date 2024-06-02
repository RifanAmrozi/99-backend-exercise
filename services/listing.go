package services

import (
	"99-backend-exercise/models"
	"99-backend-exercise/repositories"
)

func CreateListing(listing *models.Listing) error {
	return repositories.CreateListing(listing)
}
