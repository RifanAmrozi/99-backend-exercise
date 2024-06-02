package models

import "time"

type Listing struct {
	ID          int       `db:"id" json:"id"`
	UserID      int       `db:"user_id" json:"user_id"`
	Price       int       `db:"price" json:"price"`
	ListingType string    `db:"listing_type" json:"listing_type"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
