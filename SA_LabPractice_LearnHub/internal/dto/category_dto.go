package dto

import "time"

// CategoryResponse defines the payload returned for category queries.
type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateCategoryRequest defines payload for category creation.
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2,max=120"`
	Slug string `json:"slug" binding:"required,min=2,max=120"`
}
