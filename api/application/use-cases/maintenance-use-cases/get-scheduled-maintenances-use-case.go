package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
)

type GetScheduledMaintenancesUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewGetScheduledMaintenancesUseCase(maintenanceRepository repositories.MaintenanceRepository) *GetScheduledMaintenancesUseCase {
	return &GetScheduledMaintenancesUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetScheduledMaintenancesUseCase) Execute() ([]*m.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*m.MaintenanceOutputDTO)
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

func (useCase *GetScheduledMaintenancesUseCase) performGetScheduledMaintenances(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*m.MaintenanceOutputDTO, validationErrorSignal chan<- bool) {
	defer wg.Done()

	var scheduledMaintenances []*domain.Maintenance

	go func() {
		var err error
		if scheduledMaintenances, err = useCase.maintenanceRepository.GetScheduledMaintenances(); err != nil {
			errorChan <- fmt.Errorf("failed to retrieve scheduled maintenance records: %w", err)
			validationErrorSignal <- true
			return
		}

		resultChan <- maintUt.ConvertMaintenanceListToDTOs(scheduledMaintenances)
		validationErrorSignal <- false
	}()
}
