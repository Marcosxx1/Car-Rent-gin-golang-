package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CarID       string `json:"car_id"`
	Car         *[]Car
}
