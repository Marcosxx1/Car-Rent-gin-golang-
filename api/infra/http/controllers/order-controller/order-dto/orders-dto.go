package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusActive    OrderStatus = "active"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderInputDTO struct {
	UserID          string      `json:"user_id" validate:"required,notnull"`
	CarID           string      `json:"car_id" validate:"required,notnull"`
	RentalStartDate time.Time   `json:"rental_start_date" validate:"required"`
	RentalEndDate   time.Time   `json:"rental_end_date" validate:"required,gtfield=RentalStartDate"`
	TotalCost       float64     `json:"total_cost" validate:"required,gt=0,notnull"`
	OrderStatus     OrderStatus `json:"order_status" validate:"required,oneof=pending confirmed active completed cancelled"`
}

type OrderOutputDTO struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	CarID           string    `json:"car_id"`
	RentalStartDate time.Time `json:"rental_start_date"`
	RentalEndDate   time.Time `json:"rental_end_date"`
	TotalCost       float64   `json:"total_cost"`
	OrderStatus     bool      `json:"order_status"`
}

func (o *OrderInputDTO) Validate() error {
	return validate.Struct(o)
}
