package orderusecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/order-controller/order-dto"
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

	// 1. Veficiar se o usuário existe
	// 2. Validar o input
	// 3. Criar a instancia do pedido
	//newOrder := domain.createOrderInstance(input.UserID, input.CarID, input.RentalStartDate, input.RentalEndDate, input.TotalCost, input.OrderStatus)

	// 4. Se tudo certo, salvar
	/* 	createdOrder, err := useCase.orderRepository.CreateOrder(newOrder)
	   	if err != nil {
	   		return nil, err
	   	} */

	// 5. Converter para o output
	//outputDTO := orderdto.ConvertToOutputDTO(createdOrder)
	//return outputDTO, nil
	return nil, nil
}
