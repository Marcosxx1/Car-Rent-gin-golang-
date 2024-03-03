package usecases

import (
	"fmt"
	"sync"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	carutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases/car-use-case-tests/car-utils"
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
)

type PutCarUseCase struct {
	carRepository           r.CarRepository
	specificationRepository r.SpecificationRepository
}

func NewUpdateCarUseCase(
	carRepository r.CarRepository,
	specificationRepository r.SpecificationRepository) *PutCarUseCase {

	return &PutCarUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *PutCarUseCase) Execute(carID string, inputDTO *dtos.CarInputDTO) (*dtos.CarOutputDTO, error) {

	resultChan := make(chan *dtos.CarOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)
 	go carutils.PerformValidationForUpdate(&wg, errorChan, validationErrorSignal, inputDTO, useCase.carRepository)
	go useCase.performCarUpdate(&wg, resultChan, errorChan, validationErrorSignal, carID, inputDTO)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case createdCar := <-resultChan:
		return createdCar, nil
	case err := <-errorChan:
		return nil, err
	}
}

 

func (useCase *PutCarUseCase) performCarUpdate(wg *sync.WaitGroup, resultChan chan<- *dtos.CarOutputDTO, errorChan chan<- error, validationErrorSignal <-chan bool, carID string, inputDTO *dtos.CarInputDTO) {
	defer wg.Done()

	if <-validationErrorSignal {
		return
	}

	var domainSpecification []*domain.Specification

	if len(inputDTO.Specification) > 0 {
		domainSpecification = repoutils.ConvertInputSpecificationToDomainUpdate(inputDTO.Specification)
	}

	carToBeUpdated := &domain.Car{
		ID:           carID,
		Name:         inputDTO.Name,
		Description:  inputDTO.Description,
		DailyRate:    inputDTO.DailyRate,
		Available:    inputDTO.Available,
		LicensePlate: inputDTO.LicensePlate,
		FineAmount:   inputDTO.FineAmount,
		Brand:        inputDTO.Brand,
	}

	carUpdated, err := useCase.carRepository.UpdateCar(carID, carToBeUpdated)
	if err != nil {
		errorChan <- fmt.Errorf("failed to update car record: %w", err)
		return
	}

	specificationUpdated, err := useCase.specificationRepository.UpdateSpecification(carID, domainSpecification)
	if err != nil {
		errorChan <- fmt.Errorf("failed to update specification record: %w", err)
	}

	resultChan <- dtos.ConvertDomainToOutPut(carID, carUpdated, specificationUpdated)
}
