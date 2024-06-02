package repositories

import (
	"99-backend-exercise/db"
	"99-backend-exercise/models"
	"database/sql"
	"time"
)

func CreateListing(listing *models.Listing) error {
	listing.CreatedAt = time.Now()
	listing.UpdatedAt = time.Now()
	query := `INSERT INTO listings (user_id, price, listing_type, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return db.DB.QueryRow(query, listing.UserID, listing.Price, listing.ListingType, listing.CreatedAt, listing.UpdatedAt).Scan(&listing.ID)
}

func GetListings(pageNum, pageSize int, userID *int) ([]models.Listing, error) {
	var listings []models.Listing
	offset := (pageNum - 1) * pageSize
	var rows *sql.Rows
	var err error

	if userID != nil {
		query := `SELECT id, user_id, price, listing_type, created_at, updated_at 
                  FROM listings WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
		rows, err = db.DB.Query(query, *userID, pageSize, offset)
	} else {
		query := `SELECT id, user_id, price, listing_type, created_at, updated_at 
                  FROM listings ORDER BY created_at DESC LIMIT $1 OFFSET $2`
		rows, err = db.DB.Query(query, pageSize, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var listing models.Listing
		if err := rows.Scan(&listing.ID, &listing.UserID, &listing.Price, &listing.ListingType, &listing.CreatedAt, &listing.UpdatedAt); err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	return listings, nil
}
