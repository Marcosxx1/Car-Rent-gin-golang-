package orderdto

import (
	"time"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
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

type OrderInputCompleteDTO struct {
	UserID          string
	CarID           string
	RentalStartDate time.Time
	RentalEndDate   time.Time
	TotalCost       float64
	OrderStatus/* OrderStatus */ bool
}

type OrderInputPartialDTO struct {
	RentalStartDate                   time.Time `json:"rental_start_date" validate:"required"`
	RentalEndDate                     time.Time `json:"rental_end_date" validate:"required,gtfield=RentalStartDate"`
	TotalCost                         float64   `json:"total_cost" validate:"required,gt=0,notnull"`
	OrderStatus/* OrderStatus */ bool           `json:"order_status" validate:"required,oneof=pending confirmed active completed cancelled"`
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

func (o *OrderInputPartialDTO) Validate() error {
	return validate.Struct(o)
}

func ConvertToOutputDTO(newOrder *domain.Order) *OrderOutputDTO {
	return &OrderOutputDTO{
		ID:              newOrder.ID,
		UserID:          newOrder.UserID,
		CarID:           newOrder.CarID,
		RentalStartDate: newOrder.RentalStartDate,
		RentalEndDate:   newOrder.RentalEndDate,
		TotalCost:       newOrder.TotalCost,
		OrderStatus:     newOrder.OrderStatus,
	}
}
