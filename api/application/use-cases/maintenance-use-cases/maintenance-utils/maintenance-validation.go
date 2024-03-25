package maintenanceutils

import (
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

func PerformMaintenanceValidation(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, inputDTO maintenancedtos.MaintenanceInputDTO) {
	defer wg.Done()

	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
		errorChan <- err
		validationErrorSignal <- true
		return
	}

	validationErrorSignal <- false
}
