package repositories

import (
	"99-backend-exercise/db"
	"99-backend-exercise/models"
	"time"
)

func CreateListing(listing *models.Listing) error {
	listing.CreatedAt = time.Now()
	listing.UpdatedAt = time.Now()
	query := `INSERT INTO listings (user_id, price, listing_type, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return db.DB.QueryRow(query, listing.UserID, listing.Price, listing.ListingType, listing.CreatedAt, listing.UpdatedAt).Scan(&listing.ID)
}
