package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
)

type ListMaintenanceUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewListMaintenanceUseCase(maintenanceRepository repositories.MaintenanceRepository) *ListMaintenanceUseCase {
	return &ListMaintenanceUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *ListMaintenanceUseCase) Execute() ([]*m.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*m.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)
	go useCase.performListMaintenance(&wg, errorChan, resultChan, validationErrorSignal)

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
func (useCase *ListMaintenanceUseCase) performListMaintenance(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*m.MaintenanceOutputDTO, validationErrorSignal chan<- bool) {
	defer wg.Done()

	var maintenances []*domain.Maintenance

	go func() { // comentar com Hllx
		var err error
		if maintenances, err = useCase.maintenanceRepository.ListAllMaintenances(); err != nil {
			errorChan <- fmt.Errorf("failed to create maintenance record: %w", err)
			validationErrorSignal <- true
			return
		}

		resultChan <- maintUt.ConvertMaintenanceListToDTOs(maintenances)
		validationErrorSignal <- false
	}()
}
