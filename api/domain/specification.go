package domain

import "gorm.io/gorm"

type Specification struct {
	gorm.Model
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CarID       string `json:"car_id"`
	Car         *Car
}
