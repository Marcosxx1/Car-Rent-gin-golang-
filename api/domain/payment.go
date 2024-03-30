package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID            string    `json:"id"`
	OrderID       uint      `json:"order_id"`
	Order         Order     `gorm:"foreignKey:OrderID"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentStatus string    `json:"payment_status"`
}
