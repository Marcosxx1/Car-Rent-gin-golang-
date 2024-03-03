package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
)

type GetMaintenanceByStatusUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewNewGetMaintenanceByStatusUseCase(maintenanceRepository repositories.MaintenanceRepository) *GetMaintenanceByStatusUseCase {
	return &GetMaintenanceByStatusUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *GetMaintenanceByStatusUseCase) Execute(maintenance_status bool) ([]*m.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*m.MaintenanceOutputDTO)
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

func (useCase *GetMaintenanceByStatusUseCase) performGetMaintenanceByStatus(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*m.MaintenanceOutputDTO, validationErrorSignal chan<- bool, maintenance_status bool) {
	defer wg.Done()

	var maintenances []*domain.Maintenance

	go func() {
		var err error
		if maintenances, err = useCase.maintenanceRepository.GetMaintenancesByStatus(maintenance_status); err != nil {
			errorChan <- fmt.Errorf("failed to retrieve maintenance records by its status: %w", err)
			validationErrorSignal <- true
			return
		}

		resultChan <- maintUt.ConvertMaintenanceListToDTOs(maintenances)
		validationErrorSignal <- false
	}()
}
