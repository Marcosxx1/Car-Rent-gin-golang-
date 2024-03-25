package orderusecases

import (
	"fmt"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type PostOrderUseCase struct {
	userRepository        repositories.UserRepository
	maintenanceRepository repositories.MaintenanceRepository
	orderRepository       repositories.OrderRepository
}

func NewPostOrderUseCase(
	userRepository repositories.UserRepository,
	maintenanceRepository repositories.MaintenanceRepository,
	orderRepository repositories.OrderRepository) *PostOrderUseCase {
	return &PostOrderUseCase{
		userRepository:        userRepository,
		maintenanceRepository: maintenanceRepository,
		orderRepository:       orderRepository,
	}
}

func (useCase *PostOrderUseCase) Execute(input orderdto.OrderInputCompleteDTO) (*orderdto.OrderOutputDTO, error) {
	user, err := useCase.userRepository.GetById(input.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by ID %s: %w", input.UserID, err)
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %s does not exist", input.UserID)
	}

	newOrder, err := domain.CreateOrderInstance(input.UserID, input.CarID, input.RentalStartDate, input.RentalEndDate, input.TotalCost, input.OrderStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to create order instance: %w", err)
	}

	err = useCase.orderRepository.CreateOrder(newOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	outputDTO := orderdto.ConvertToOutputDTO(newOrder)
	return outputDTO, nil
}
