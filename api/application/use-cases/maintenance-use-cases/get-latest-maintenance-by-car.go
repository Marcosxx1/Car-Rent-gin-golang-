package maintenanceusecases

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type GetLatestMaintenanceByCar struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewGetLatestMaintenanceByCarIDUseCase(maintenanceRepository repositories.MaintenanceRepository) *GetLatestMaintenanceByCar {
	return &GetLatestMaintenanceByCar{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetLatestMaintenanceByCar) Execute(carID string) (*maintenancedtos.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan *maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)
	go useCase.performGetLatestMaintenanceByCar(&wg, errorChan, resultChan, validationErrorSignal, carID)

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

func (useCase *GetLatestMaintenanceByCar) performGetLatestMaintenanceByCar(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- *maintenancedtos.MaintenanceOutputDTO, validationErrorSignal chan<- bool, carID string) {
	defer wg.Done()

	var maintenances *domain.Maintenance
	var err error

	go func() {
		if maintenances, err = useCase.maintenanceRepository.GetLatestMaintenanceByCar(carID); err != nil {
			errorChan <- fmt.Errorf("failed to retrieve maintenance records by carID: %w", err)
			validationErrorSignal <- true
			return
		}

		resultChan <- maintenanceutils.ConvertToOutputDTO(maintenances)
		validationErrorSignal <- false
	}()
}
