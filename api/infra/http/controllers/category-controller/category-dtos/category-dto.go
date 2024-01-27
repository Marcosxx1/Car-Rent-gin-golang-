package categorydtos

import "time"

// CategoryInputDTO represents the data expected when creating or updating a category.
type CategoryInputDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// CategoryOutputDTO represents the data returned when retrieving a category.
type CategoryOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
