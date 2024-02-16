package usecases

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
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

func (useCase *PostCarUseCase) ExecuteConcurrently(inputDTO *dtos.CarInputDTO) (*dtos.CarOutputDTO, error) {
	carID := xid.New().String()

	resultChan := make(chan *dtos.CarOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool) // Um sinal para caso tenhamos um erro > true 

	var wg sync.WaitGroup

	wg.Add(2)																		// passamos ele nas goroutines
	go useCase.performValidation(&wg, errorChan, validationErrorSignal, inputDTO)
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

func (useCase *PostCarUseCase) performValidation(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, inputDTO *dtos.CarInputDTO) {
	defer wg.Done()

	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
			errorChan <- err
			validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
			return
	}

	existCar, _ := useCase.carRepository.FindCarByLicensePlate(inputDTO.LicensePlate)

	if existCar != nil {
			errorChan <- errors.New("a car with the same license plate already exists")
			validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
			return
	}

	validationErrorSignal <- false // Sinaliza que a validação foi bem-sucedida
}

func (useCase *PostCarUseCase) performCarCreation(wg *sync.WaitGroup, resultChan chan<- *dtos.CarOutputDTO, errorChan chan<- error, validationErrorSignal <-chan bool, carID string, inputDTO *dtos.CarInputDTO) {
	defer wg.Done()

	if <-validationErrorSignal { // Verifica se houve erro de validação
			// Não prossegue com a criação da instância se houver erro de validação
			return
	}

	specifications := repoutils.ConvertSpecificationToDomainCreate(inputDTO.Specification, carID)

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

	resultChan <- dtos.ConvertToOutputDTO(carID, inputDTO)
}
