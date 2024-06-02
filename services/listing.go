package services

import (
	"99-backend-exercise/models"
	"99-backend-exercise/repositories"
)

func CreateListing(listing *models.Listing) error {
	return repositories.CreateListing(listing)
}
func GetListings(pageNum, pageSize int, userID *int) ([]models.Listing, error) {
	return repositories.GetListings(pageNum, pageSize, userID)
}
