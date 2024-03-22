package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
)

type PatchMaintenanceUseCase struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewPatchMaintenanceUseCase(
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) *PatchMaintenanceUseCase {
	return &PatchMaintenanceUseCase{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *PatchMaintenanceUseCase) Execute(maintenanceID string, inputDTO maintenancedtos.MaintenanceInputDTO) (*maintenancedtos.MaintenanceOutputDTO, error) {
	// Create channels for the result and errors
	resultChan := make(chan *maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	var wg sync.WaitGroup

	// Add counters for the goroutines
	wg.Add(2)
	go maintenanceutils.PerformMaintenanceValidation(&wg, errorChan, validationErrorSignal, inputDTO)
	go useCase.performMaintenanceUpdate(&wg, errorChan, validationErrorSignal, resultChan, maintenanceID, inputDTO)

	// Add counter for the goroutine to close channels
	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Wait for the results or errors
	select {
	case updatedMaintenance := <-resultChan:
		return updatedMaintenance, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *PatchMaintenanceUseCase) performMaintenanceUpdate(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, resultChan chan<- *maintenancedtos.MaintenanceOutputDTO, maintenanceID string, inputDTO maintenancedtos.MaintenanceInputDTO) {
	defer wg.Done()

	existingMaintenance, err := useCase.maintenanceRepository.GetMaintenanceByID(maintenanceID)
	if err != nil {
		errorChan <- fmt.Errorf("failed to retrieve existing maintenance record: %w", err)
		validationErrorSignal <- true
		return
	}
	if existingMaintenance == nil {
		errorChan <- fmt.Errorf("maintenance not found")
		validationErrorSignal <- true
		return
	}

	maintenanceToUpdate, err := domain.CreateMaintenanceInstance(inputDTO.CarID, inputDTO)
	if err != nil {
		errorChan <- err
		validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
		return
	}

	go func() {
		if err := useCase.maintenanceRepository.UpdateMaintenance(maintenanceToUpdate, maintenanceID); err != nil {
			errorChan <- fmt.Errorf("failed to update maintenance record: %w", err)
			validationErrorSignal <- true
			return
		}
	}()

	resultChan <- maintenanceutils.ConvertToOutputDTO(maintenanceToUpdate)
	validationErrorSignal <- false
}
