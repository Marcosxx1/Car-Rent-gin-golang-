package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
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

func (useCase *GetMaintenancesByDateRangeUseCase) Execute(startDate, endDate string) ([]*m.MaintenanceOutputDTO, error) {
	resultChan := make(chan []*m.MaintenanceOutputDTO)
	errorChan := make(chan error)

	var wg sync.WaitGroup

	wg.Add(2)
	go maintUt.ValidateDateRange(&wg, errorChan, startDate, endDate)
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

func (useCase *GetMaintenancesByDateRangeUseCase) performMaintenanceRetrieval(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*m.MaintenanceOutputDTO, startDate, endDate string) {
	defer wg.Done()

	maintenances, err := useCase.maintenanceRepository.GetMaintenancesByDateRange(startDate, endDate)
	if err != nil {
		errorChan <- fmt.Errorf("failed to retrieve maintenance records: %w", err)
		return
	}

	outputDTOs := make([]*m.MaintenanceOutputDTO, len(maintenances))
	for i, maintenance := range maintenances {
		outputDTOs[i] = maintUt.ConvertToOutputDTO(maintenance)
	}

	resultChan <- outputDTOs
}
