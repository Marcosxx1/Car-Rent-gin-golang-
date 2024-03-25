package categorydtos

import "time"

type CategoryInputDTO struct {
	Name        string `json:"name" validate:"required" example:"'name is required'"`
	Description string `json:"description" validate:"required" example:"'description is required'"`
}


type CategoryOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
