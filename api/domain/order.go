package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	UserID          string
	User            User `gorm:"foreignKey:user_id"`
	CarID           string
	Car             Car `gorm:"foreignKey:car_id"`
	RentalStartDate time.Time
	RentalEndDate   time.Time
	TotalCost       float64
	OrderStatus     bool
}
