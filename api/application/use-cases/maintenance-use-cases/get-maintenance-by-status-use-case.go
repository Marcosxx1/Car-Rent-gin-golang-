package maintenanceusecases

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain/enums"
)

type GetMaintenanceByStatusUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewGetMaintenanceByStatusUseCase(maintenanceRepository repositories.MaintenanceRepository) *GetMaintenanceByStatusUseCase {
	return &GetMaintenanceByStatusUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetMaintenanceByStatusUseCase) Execute(maintenance_status enums.MaintenanceStatus) ([]*maintenancedtos.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)
	go useCase.performGetMaintenanceByStatus(&wg, errorChan, resultChan, validationErrorSignal, maintenance_status)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case maintenanceList := <-resultChan:
		return maintenanceList, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *GetMaintenanceByStatusUseCase) performGetMaintenanceByStatus(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*maintenancedtos.MaintenanceOutputDTO, validationErrorSignal chan<- bool, maintenance_status enums.MaintenanceStatus) {
	defer wg.Done()

	var maintenances []*domain.Maintenance

	go func() {
		var err error
		if maintenances, err = useCase.maintenanceRepository.GetMaintenancesByStatus(maintenance_status); err != nil {
			errorChan <- fmt.Errorf("failed to retrieve maintenance records by its status: %w", err)
			validationErrorSignal <- true
			return
		}

		if len(maintenances) == 0 {
			resultChan <- []*maintenancedtos.MaintenanceOutputDTO{}
			validationErrorSignal <- false
			return
		}

		resultChan <- maintenanceutils.ConvertMaintenanceListToDTOs(maintenances)
		validationErrorSignal <- false
	}()
}
