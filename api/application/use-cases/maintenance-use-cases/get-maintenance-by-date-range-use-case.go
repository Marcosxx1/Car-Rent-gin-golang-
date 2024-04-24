package maintenanceusecases

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/maintenance-utils"
)

type GetMaintenancesByDateRangeUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewGetMaintenancesByDateRangeUseCase(
	maintenanceRepository repositories.MaintenanceRepository) *GetMaintenancesByDateRangeUseCase {
	return &GetMaintenancesByDateRangeUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetMaintenancesByDateRangeUseCase) Execute(startDate, endDate string) ([]*maintenancedtos.MaintenanceOutputDTO, error) {
	resultChan := make(chan []*maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)

	var wg sync.WaitGroup

	wg.Add(2)
	go maintenanceutils.ValidateDateRange(&wg, errorChan, startDate, endDate)
	go useCase.performMaintenanceRetrieval(&wg, errorChan, resultChan, startDate, endDate)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case maintenances := <-resultChan:
		return maintenances, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *GetMaintenancesByDateRangeUseCase) performMaintenanceRetrieval(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*maintenancedtos.MaintenanceOutputDTO, startDate, endDate string) {
	defer wg.Done()

	maintenances, err := useCase.maintenanceRepository.GetMaintenancesByDateRange(startDate, endDate)
	if err != nil {
		errorChan <- fmt.Errorf("failed to retrieve maintenance records: %w", err)
		return
	}

	outputDTOs := make([]*maintenancedtos.MaintenanceOutputDTO, len(maintenances))
	for i, maintenance := range maintenances {
		outputDTOs[i] = maintenanceutils.ConvertToOutputDTO(maintenance)
	}

	resultChan <- outputDTOs
}
