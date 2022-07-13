package models

type ListingImage struct {
	ID        string `gorm:"primaryKey"`
	ListingID string
	ImageURL  string
}
