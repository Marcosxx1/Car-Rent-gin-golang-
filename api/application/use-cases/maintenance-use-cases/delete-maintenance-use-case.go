package maintenanceusecases

import (
	"fmt"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/maintenance-utils"
)

type DeleteMaintenanceUseCase struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewDeleteMaintenanceUseCase(
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) *DeleteMaintenanceUseCase {
	return &DeleteMaintenanceUseCase{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *DeleteMaintenanceUseCase) Execute(maintenanceID string) error {
	resultChan := make(chan *maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)

	go useCase.performMaintenanceDeletion(errorChan, resultChan, maintenanceID)

	select {
	case err := <-errorChan:
		return err
	case <-resultChan:
		return nil
	}
}

func (useCase *DeleteMaintenanceUseCase) performMaintenanceDeletion(errorChan chan<- error, resultChan chan<- *maintenancedtos.MaintenanceOutputDTO, maintenanceID string) {
	defer close(resultChan)
	defer close(errorChan)

	existingMaintenance, err := useCase.maintenanceRepository.GetMaintenanceByID(maintenanceID)
	if err != nil {
		errorChan <- fmt.Errorf("failed to retrieve existing maintenance record: %w", err)
		return
	}

	if err := useCase.maintenanceRepository.DeleteMaintenance(maintenanceID); err != nil {
		errorChan <- fmt.Errorf("failed to delete maintenance record: %w", err)
		return
	}

	resultChan <- maintenanceutils.ConvertToOutputDTO(existingMaintenance)
}
