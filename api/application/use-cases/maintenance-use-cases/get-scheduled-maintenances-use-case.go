package maintenanceusecases

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type GetScheduledMaintenancesUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewGetScheduledMaintenancesUseCase(maintenanceRepository repositories.MaintenanceRepository) *GetScheduledMaintenancesUseCase {
	return &GetScheduledMaintenancesUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetScheduledMaintenancesUseCase) Execute() ([]*maintenancedtos.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)
	go useCase.performGetScheduledMaintenances(&wg, errorChan, resultChan, validationErrorSignal)

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

func (useCase *GetScheduledMaintenancesUseCase) performGetScheduledMaintenances(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*maintenancedtos.MaintenanceOutputDTO, validationErrorSignal chan<- bool) {
	defer wg.Done()

	var scheduledMaintenances []*domain.Maintenance

	go func() {
		var err error
		if scheduledMaintenances, err = useCase.maintenanceRepository.GetScheduledMaintenances(); err != nil {
			errorChan <- fmt.Errorf("failed to retrieve scheduled maintenance records: %w", err)
			validationErrorSignal <- true
			return
		}

		resultChan <- maintenanceutils.ConvertMaintenanceListToDTOs(scheduledMaintenances)
		validationErrorSignal <- false
	}()
}
