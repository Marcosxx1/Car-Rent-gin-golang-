package orderusecases

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type DeleteOrderUseCase struct {
	orderRepository repositories.OrderRepository
}

func NewDeleteOrderUseCase(orderRepository repositories.OrderRepository) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (useCase *DeleteOrderUseCase) Execute(orderId string) error {
	err := useCase.performOrderDeletion(orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)
	}
	return nil
}

func (useCase *DeleteOrderUseCase) performOrderDeletion(orderId string) error {
	if orderId == "" {
		return fmt.Errorf("orderId cannot be empty")
	}


	err := useCase.orderRepository.DeleteOrder(orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)
	}

	return nil
}
