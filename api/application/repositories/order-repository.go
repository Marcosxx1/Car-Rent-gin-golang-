package repositories

import (
	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type OrderRepository interface {
	// CreateOrder creates a new order in the database
	CreateOrder(order *domain.Order) error

	// GetOrderByID retrieves an order by its ID
	GetOrdersByOptions(options *orderdto.OrderOutputDTO) ([]*domain.Order, error)
	// UpdateOrder updates an existing order in the database
	UpdateOrder(id string, order *domain.Order) error
	// DeleteOrder deletes an order from the database
	DeleteOrder(orderID string) error

	// DeleteOrder deletes an order from the database
	/* 	DeleteOrder(orderID string) error

	   	// ListOrdersByUserID retrieves a list of orders for a given user ID
	   	ListOrdersByUserID(userID string) ([]*domain.Order, error)

	   	// ListOrdersByCarID retrieves a list of orders for a given car ID
	   	ListOrdersByCarID(carID string) ([]*domain.Order, error) */
}
