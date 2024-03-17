package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
)

type GetMaintenancesByCarIDUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewGetMaintenancesByCarIDUseCase(maintenanceRepository repositories.MaintenanceRepository) *GetMaintenancesByCarIDUseCase {
	return &GetMaintenancesByCarIDUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetMaintenancesByCarIDUseCase) Execute(carID string) ([]*m.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*m.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)
	go useCase.performGetMaintenancesByCarID(&wg, errorChan, resultChan, validationErrorSignal, carID)

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

func (useCase *GetMaintenancesByCarIDUseCase) performGetMaintenancesByCarID(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*m.MaintenanceOutputDTO, validationErrorSignal chan<- bool, carID string) {
	defer wg.Done()

	var maintenances []*domain.Maintenance

	go func() {
		var err error
		if maintenances, err = useCase.maintenanceRepository.GetMaintenancesByCarID(carID); err != nil {
			errorChan <- fmt.Errorf("failed to retrieve maintenance records by carID: %w", err)
			validationErrorSignal <- true
			return
		}

		resultChan <- maintUt.ConvertMaintenanceListToDTOs(maintenances)
		validationErrorSignal <- false
	}()
}
