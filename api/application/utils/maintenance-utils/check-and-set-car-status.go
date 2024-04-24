package maintenanceutils

import (
	"fmt"
	"sync"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

func CheckAndSetStatus(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, resultChan chan<- *maintenancedtos.MaintenanceOutputDTO, carID string, carRepository repositories.CarRepository) {
	defer wg.Done()

	car, err := carRepository.FindCarById(carID)
	if err != nil {
		errorChan <- err
		validationErrorSignal <- true
		return
	}

	if !car.Available {
		errorChan <- fmt.Errorf("car with id %s is not available or is in maintenance", carID)
		validationErrorSignal <- true
		return
	}

	if err := carRepository.AlterCarStatus(carID, false); err != nil {
		errorChan <- err
		validationErrorSignal <- true
		return
	}

	validationErrorSignal <- false
}
