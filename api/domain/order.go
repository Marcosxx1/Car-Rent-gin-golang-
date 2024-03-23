package domain

import (
	"time"

	"github.com/rs/xid"
)

// Order represents an order entity.
type Order struct {
	ID              string // `gorm:"primaryKey"`
	UserID          string
	User            User // `gorm:"foreignKey:user_id"`
	CarID           string
	Car             Car // `gorm:"foreignKey:car_id"`
	RentalStartDate time.Time
	RentalEndDate   time.Time
	TotalCost       float64
	OrderStatus     bool
}

// CreateOrderInstance creates a new instance of the Order struct.
func CreateOrderInstance(userID, carID string, rentalStartDate, rentalEndDate time.Time, totalCost float64, orderStatus bool) (*Order, error) {
	orderID := generateOrderID()

	return &Order{
		ID:              orderID,
		UserID:          userID,
		CarID:           carID,
		RentalStartDate: rentalStartDate,
		RentalEndDate:   rentalEndDate,
		TotalCost:       totalCost,
		OrderStatus:     orderStatus,
	}, nil
}

// generateOrderID generates a unique order ID.
func generateOrderID() string {
	id := xid.New().String()

	return id
}
