package domain

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID            string           `json:"id"`
	Name          string           `json:"name" validate:"required"`
	Description   string           `json:"description" validate:"required"`
	DailyRate     float64          `json:"daily_rate" validate:"gte=0"`
	Available     bool             `json:"available"`
	LicensePlate  string           `json:"license_plate" validate:"required"`
	FineAmount    float64          `json:"fine_amount" validate:"gte=0"`
	Brand         string           `json:"brand" validate:"required"`
	CategoryID    string           `json:"category_id"`
	Specification []*Specification `json:"specifications"` // One car HAS MANY specifications
	Maintenances  []*Maintenance   `gorm:"many2many:car_maintenances;"`
}
