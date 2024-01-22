package dtos

import "time"


 type CarInputDTO struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	DailyRate    float64 `json:"daily_rate" validate:"required"`
	Available    bool    `json:"available" validate:"required"`
	LicensePlate string  `json:"license_plate" validate:"required"`
	FineAmount   float64 `json:"fine_amount" validate:"required"`
	Brand        string  `json:"brand" validate:"required"`
	CategoryId   string  `json:"category_id" validate:"required"`
 }

 type CarOutputDTO struct {
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


/* type CarCommonDTO struct {
    Name         string  `json:"name" validate:"required"`
    Description  string  `json:"description" validate:"required"`
    DailyRate    float64 `json:"daily_rate" validate:"required"`
    Available    bool    `json:"available" validate:"required"`
    LicensePlate string  `json:"license_plate" validate:"required"`
    FineAmount   float64 `json:"fine_amount" validate:"required"`
    Brand        string  `json:"brand" validate:"required"`
    CategoryId   string  `json:"category_id" validate:"required"`
}

type CarInputDTO struct {
    CarCommonDTO
}

type CarOutputDTO struct {
    CarCommonDTO
    Id        string    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
}
 */