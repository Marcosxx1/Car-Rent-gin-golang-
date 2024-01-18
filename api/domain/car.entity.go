package domain

import "time"

type Car struct {
	Id           string    `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	DailyRate    float64   `json:"daily_rate" validate:"gte=0"`
	Available    bool      `json:"available"`
	LicensePlate string    `json:"license_plate" validate:"required"`
	FineAmount   float64   `json:"fine_amount" validate:"gte=0"`
	Brand        string    `json:"brand" validate:"required"`
	CategoryId   string    `json:"category_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
}
