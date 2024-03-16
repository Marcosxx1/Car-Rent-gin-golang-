package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type OrderRepository interface {
	// CreateOrder creates a new order in the database
	CreateOrder(order domain.Order) error

	/* 	// GetOrderByID retrieves an order by its ID
	   	GetOrderByID(orderID string) (*models.Order, error)

	   	// UpdateOrder updates an existing order in the database
	   	UpdateOrder(order *models.Order) error

	   	// DeleteOrder deletes an order from the database
	   	DeleteOrder(orderID string) error

	   	// ListOrdersByUserID retrieves a list of orders for a given user ID
	   	ListOrdersByUserID(userID string) ([]*models.Order, error)

	   	// ListOrdersByCarID retrieves a list of orders for a given car ID
	   	ListOrdersByCarID(carID string) ([]*models.Order, error) */
}
