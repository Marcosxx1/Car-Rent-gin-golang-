package maintenanceutils

import (
	"sync"

	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

func PerformMaintenanceValidation(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, inputDTO m.MaintenanceInputDTO) {
	defer wg.Done()

	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
		errorChan <- err
		validationErrorSignal <- true
		return
	}

	validationErrorSignal <- false
}
