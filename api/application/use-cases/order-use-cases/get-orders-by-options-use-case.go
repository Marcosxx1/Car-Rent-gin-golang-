package orderusecases

import (
	"fmt"
	"sync"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type GetOrderByQueryUseCase struct {
	userRepository        repositories.UserRepository
	maintenanceRepository repositories.MaintenanceRepository
	orderRepository       repositories.OrderRepository
}

func NewGetOrderByQueryUseCase(
	userRepository repositories.UserRepository,
	maintenanceRepository repositories.MaintenanceRepository,
	orderRepository repositories.OrderRepository) *GetOrderByQueryUseCase {
	return &GetOrderByQueryUseCase{
		userRepository:        userRepository,
		maintenanceRepository: maintenanceRepository,
		orderRepository:       orderRepository,
	}
}

func (useCase *GetOrderByQueryUseCase) Execute(options *orderdto.Options) ([]*orderdto.OrderOutputDTO, error) {
	var wg sync.WaitGroup
	resultChan := make(chan []*orderdto.OrderOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)

	go useCase.performGetOrderByOptions(&wg, errorChan, resultChan, validationErrorSignal, options)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case orderList := <-resultChan:
		return orderList, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *GetOrderByQueryUseCase) performGetOrderByOptions(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*orderdto.OrderOutputDTO, validationErrorSignal chan<- bool, options *orderdto.Options) {
	defer wg.Done()

	orders, err := useCase.orderRepository.GetOrdersByOptions(options)
	if err != nil {
		errorChan <- fmt.Errorf("failed to retrieve order details: %w", err)
		validationErrorSignal <- true
		return
	}

	if len(orders) == 0 {
		resultChan <- []*orderdto.OrderOutputDTO{}
		validationErrorSignal <- false
		return
	}

	resultChan <- orderdto.ConvertToOutputDTOList(orders)
	validationErrorSignal <- false
}
