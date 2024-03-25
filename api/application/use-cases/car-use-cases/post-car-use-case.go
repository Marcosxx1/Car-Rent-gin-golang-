package carusecases

import (
	"fmt"
	"sync"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	carutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases/car-use-case-tests/car-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/rs/xid"
)

type PostCarUseCase struct {
	carRepository           repositories.CarRepository
	specificationRepository repositories.SpecificationRepository
}

func NewPostCarUseCase(
	carRepository repositories.CarRepository,
	specificationRepository repositories.SpecificationRepository) *PostCarUseCase {
	return &PostCarUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *PostCarUseCase) ExecuteConcurrently(inputDTO *cardtos.CarInputDTO) (*cardtos.CarOutputDTO, error) {
	carID := xid.New().String()

	resultChan := make(chan *cardtos.CarOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool) // Um sinal para caso tenhamos um erro > true

	var wg sync.WaitGroup

	wg.Add(2) // passamos ele nas goroutines
	go carutils.PerformValidation(&wg, errorChan, validationErrorSignal, inputDTO, useCase.carRepository)
	go useCase.performCarCreation(&wg, resultChan, errorChan, validationErrorSignal, carID, inputDTO)

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

func (useCase *PostCarUseCase) performCarCreation(wg *sync.WaitGroup, resultChan chan<- *cardtos.CarOutputDTO, errorChan chan<- error, validationErrorSignal <-chan bool, carID string, inputDTO *cardtos.CarInputDTO) {
	defer wg.Done()

	if <-validationErrorSignal { // Verifica se houve erro de validação
		// Não prossegue com a criação da instância se houver erro de validação
		return
	}

	var specifications []*domain.Specification
	go func() {
		specifications = utils.ConvertSpecificationToDomainCreate(inputDTO.Specification, carID)
	}()

	newCar := &domain.Car{
		ID:           carID,
		Name:         inputDTO.Name,
		Description:  inputDTO.Description,
		DailyRate:    inputDTO.DailyRate,
		Available:    inputDTO.Available,
		LicensePlate: inputDTO.LicensePlate,
		FineAmount:   inputDTO.FineAmount,
		Brand:        inputDTO.Brand,
		CategoryID:   inputDTO.CategoryID,
	}

	if err := useCase.carRepository.RegisterCar(newCar); err != nil {
		errorChan <- fmt.Errorf("failed to create car record: %w", err)
		return
	}

	if err := useCase.specificationRepository.PostMultipleSpecifications(specifications); err != nil {
		errorChan <- fmt.Errorf("failed to create specification record: %w", err)
		return
	}

	resultChan <- cardtos.ConvertToOutputDTO(carID, inputDTO)
}
