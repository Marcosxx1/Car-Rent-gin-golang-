package dtos

import "time"

type CarDto struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	DailyRate    float64 `json:"daily_rate"`
	Available    bool    `json:"available"`
	LicensePlate string  `json:"license_plate"`
	FineAmount   float64 `json:"fine_amount"`
	Brand        string  `json:"brand"`
	CategoryId   string  `json:"category_id"`
	CreatedAt    time.Time  `json:"created_at"`
}
