package maintenanceusecases

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type ListMaintenanceUseCase struct {
	maintenanceRepository repositories.MaintenanceRepository
}

func NewListMaintenanceUseCase(maintenanceRepository repositories.MaintenanceRepository) *ListMaintenanceUseCase {
	return &ListMaintenanceUseCase{
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *ListMaintenanceUseCase) Execute(page, pageSize int) ([]*maintenancedtos.MaintenanceOutputDTO, error) {
	var wg sync.WaitGroup

	resultChan := make(chan []*maintenancedtos.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	wg.Add(1)
	go useCase.performListMaintenance(&wg, errorChan, resultChan, validationErrorSignal, page, pageSize)

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
func (useCase *ListMaintenanceUseCase) performListMaintenance(wg *sync.WaitGroup, errorChan chan<- error, resultChan chan<- []*maintenancedtos.MaintenanceOutputDTO, validationErrorSignal chan<- bool, page, pageSize int) {
	defer wg.Done()

	var maintenances []*domain.Maintenance

	go func() { // comentar com Hllx
		var err error
		if maintenances, err = useCase.maintenanceRepository.ListAllMaintenances(page, pageSize); err != nil {
			errorChan <- fmt.Errorf("failed to create maintenance record: %w", err)
			validationErrorSignal <- true
			return
		}
		/* 		for _, maintenancedtos := range maintenances {
			fmt.Printf("%+v\n", maintenancedtos)
		} */
		resultChan <- maintenanceutils.ConvertMaintenanceListToDTOs(maintenances)
		validationErrorSignal <- false
	}()
}
