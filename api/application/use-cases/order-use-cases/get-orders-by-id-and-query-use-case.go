package orderusecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/order-controller/order-dto"
)

type GetOrderByIdAndQueryUseCase struct {
	userRepository        repositories.UserRepository
	maintenanceRepository repositories.MaintenanceRepository
	orderRepository       repositories.OrderRepository
}

func NewGetOrderByIdAndQueryUseCase(
	userRepository repositories.UserRepository,
	maintenanceRepository repositories.MaintenanceRepository,
	orderRepository repositories.OrderRepository) *GetOrderByIdAndQueryUseCase {
	return &GetOrderByIdAndQueryUseCase{
		userRepository:        userRepository,
		maintenanceRepository: maintenanceRepository,
		orderRepository:       orderRepository,
	}
}

func (useCase *GetOrderByIdAndQueryUseCase) Execute(orderId string) ([]*orderdto.OrderOutputDTO, error) {
	return nil, nil
}
