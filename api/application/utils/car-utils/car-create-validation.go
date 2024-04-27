package carutils

import (
	"errors"
	"sync"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

func PerformValidation(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, inputDTO *cardtos.CarInputDTO, carRepository repositories.CarRepository) {
	defer wg.Done()

	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
		errorChan <- err
		validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
		return
	}

	existCar, _ := carRepository.FindCarByLicensePlate(inputDTO.LicensePlate)

	if existCar != nil {
		errorChan <- errors.New("a car with the same license plate already exists")
		validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
		return
	}

	validationErrorSignal <- false // Sinaliza que a validação foi bem-sucedida
}
