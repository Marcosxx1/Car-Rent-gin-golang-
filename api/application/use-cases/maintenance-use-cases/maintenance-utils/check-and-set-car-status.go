package maintenanceutils

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
)

func CheckAndSetStatus(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, resultChan chan<- *m.MaintenanceOutputDTO, carID string, carRepository repositories.CarRepository) {
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
