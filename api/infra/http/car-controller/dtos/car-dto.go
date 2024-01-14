package dtos

import "time"

type CarDto struct {
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	DailyRate    float64 `json:"daily_rate" binding:"required"`
	Available    bool    `json:"available" binding:"required"`
	LicensePlate string  `json:"license_plate" binding:"required"`
	FineAmount   float64 `json:"fine_amount" binding:"required"`
	Brand        string  `json:"brand" binding:"required"`
	CategoryId   string  `json:"category_id" binding:"required"`
	CreatedAt    time.Time  `json:"created_at"`
}
