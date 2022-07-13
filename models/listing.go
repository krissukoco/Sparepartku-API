package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
	ToID      string `json:"to_id"`
	CreatedAt string `json:"created_at"`
}

type Review struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	ListingID    string    `json:"-"`
	UserID       string    `json:"user_id"`
	Star         uint      `json:"star"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"created_at"`
	IsNameHidden bool      `json:"is_name_hidden"`
}

type Listing struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	SellerID    string         `json:"seller_id"`
	ImageURLs   []string       `json:"image_urls" gorm:"-"`
	ViewCount   uint           `json:"view_count"`
	Categories  []string       `json:"categories" gorm:"-"` // max. 2 categories
	Comments    []Comment      `json:"comments" gorm:"-"`
	Reviews     []Review       `json:"reviews" gorm:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
