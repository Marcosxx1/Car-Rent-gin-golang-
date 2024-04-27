package maintenanceusecases

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/maintenance-utils"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type PostMaintenanceUseCase struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewPostMaintenanceUseCase(
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) *PostMaintenanceUseCase {
	return &PostMaintenanceUseCase{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *PostMaintenanceUseCase) Execute(carID string, inputDTO maintenancedtos.MaintenanceInputDTO) (*maintenancedtos.MaintenanceOutputDTO, error) {

	// Criamos channels para o resultado e os erros
	resultChan := make(chan *maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool) // Um sinal para caso tenhamos um erro > true

	var wg sync.WaitGroup

	// Adiciona contadores para as goroutines
	wg.Add(3)
	go maintenanceutils.PerformMaintenanceValidation(&wg, errorChan, validationErrorSignal, inputDTO)
	go maintenanceutils.CheckAndSetStatus(&wg, errorChan, validationErrorSignal, resultChan, carID, useCase.carRepository)
	go useCase.performMaintenanceCreation(&wg, errorChan, validationErrorSignal, resultChan, carID, inputDTO)

	// Adiciona contador para a goroutine de fechamento de canais
	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Esperamos pelos resultados ou erros
	select {
	case createdMaintenance := <-resultChan:
		return createdMaintenance, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *PostMaintenanceUseCase) performMaintenanceCreation(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, resultChan chan<- *maintenancedtos.MaintenanceOutputDTO, carID string, inputDTO maintenancedtos.MaintenanceInputDTO) {
	defer wg.Done()

	newMaintenance, err := domain.CreateMaintenanceInstance(carID, inputDTO)
	if err != nil {
		errorChan <- err
		validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
		return
	}

	// Inicia uma goroutine para criar a manutenção
	go func() {
		if err := useCase.maintenanceRepository.CreateMaintenance(newMaintenance); err != nil {
			errorChan <- fmt.Errorf("failed to create maintenance record: %w", err)
			validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
			return
		}
	}()

	resultChan <- maintenanceutils.ConvertToOutputDTO(newMaintenance)
	validationErrorSignal <- false
}
