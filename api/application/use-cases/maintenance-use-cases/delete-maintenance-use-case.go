package maintenanceusecases

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
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
	resultChan := make(chan *m.MaintenanceOutputDTO)
	errorChan := make(chan error)

	go useCase.performMaintenanceDeletion(errorChan, resultChan, maintenanceID)

	select {
	case err := <-errorChan:
		return err
	case <-resultChan:
		return nil
	}
}

func (useCase *DeleteMaintenanceUseCase) performMaintenanceDeletion(errorChan chan<- error, resultChan chan<- *m.MaintenanceOutputDTO, maintenanceID string) {
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

	resultChan <- maintUt.ConvertToOutputDTO(existingMaintenance)
}
