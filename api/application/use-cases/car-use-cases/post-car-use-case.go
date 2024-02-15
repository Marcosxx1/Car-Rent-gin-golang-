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
	// Generate a new car ID
	carID := xid.New().String()

	// Create channels for results and errors
	resultChan := make(chan *dtos.CarOutputDTO)
	errorChan := make(chan error)
	doneChan := make(chan struct{}) // Channel to signal completion or error

	// Create a waitgroup to wait for the completion of goroutines
	var wg sync.WaitGroup

	// Create a Once to ensure that channel closing occurs only once
	var once sync.Once
	closeChannels := func() {
		once.Do(func() {
			close(resultChan)
			close(errorChan)
			close(doneChan)
		})
	}

	// Create goroutines for each task
	wg.Add(2)
	go useCase.performValidation(&wg, errorChan, doneChan, inputDTO)
	go useCase.performCarCreation(&wg, resultChan, errorChan, carID, inputDTO)

	// Iniciamos a goroutine para fechar os canais e aguardar todas as tarefas terminarem
	wg.Add(1) // adicionamos uma goroutine que aguarda e garante as conclusões das duas goroutines acima antes de chamar o closeChannels
	go func() {
		defer wg.Done() // executado por último
		wg.Wait()       // Aguarda a conclusão das goroutines de validação e criação
		closeChannels() // fecha!
	}()

	// Wait for results or errors
	select {
	case <-doneChan: // Check if an error has occurred
		return nil, errors.New("validation or creation error")
	case createdCar := <-resultChan:
		return createdCar, nil
	case err := <-errorChan:
		return nil, err
	}
}


func (useCase *PostCarUseCase) performValidation(wg *sync.WaitGroup, errorChan chan<- error, inputDTO *dtos.CarInputDTO) {
	defer wg.Done()

	// Validate the input structure
	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
		errorChan <- err
		return
	}

	// Check if a car with the same license plate already exists
	existCar, _ := useCase.carRepository.FindCarByLicensePlate(inputDTO.LicensePlate)
	/* 	if err != nil {
		errorChan <- err
		return
	} */
	if existCar != nil {
		// A car with the same license plate already exists
		errorChan <- errors.New("a car with the same license plate already exists")
		return
	}
}

func (useCase *PostCarUseCase) performCarCreation(wg *sync.WaitGroup, resultChan chan<- *dtos.CarOutputDTO, errorChan chan<- error, carID string, inputDTO *dtos.CarInputDTO) {
	defer wg.Done()

	// Convert specification input to domain
	specifications := repoutils.ConvertSpecificationToDomainCreate(inputDTO.Specification, carID)

	// Create a new car instance
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

	// Register the new car
	if err := useCase.carRepository.RegisterCar(newCar); err != nil {
		errorChan <- fmt.Errorf("failed to create car record: %w", err)
		return
	}

	// Post multiple specifications
	if err := useCase.specificationRepository.PostMultipleSpecifications(specifications); err != nil {
		errorChan <- fmt.Errorf("failed to create specification record: %w", err)
		return
	}

	resultChan <- convertToOutputDTO(carID, inputDTO)

}

func convertToOutputDTO(carID string, inputDTO *dtos.CarInputDTO) *dtos.CarOutputDTO {
	// Assuming you have a function to convert specifications from inputDTO to outputDTO
	specificaDomain := repoutils.ConvertSpecificationToDomainCreate(inputDTO.Specification, carID)
	specificationOutPut := repoutils.ConvertSpecificationToDTO(specificaDomain)
	// Create and return the output DTO
	return &dtos.CarOutputDTO{
		ID:            carID,
		Name:          inputDTO.Name,
		Description:   inputDTO.Description,
		DailyRate:     inputDTO.DailyRate,
		Available:     inputDTO.Available,
		LicensePlate:  inputDTO.LicensePlate,
		FineAmount:    inputDTO.FineAmount,
		Brand:         inputDTO.Brand,
		CategoryID:    inputDTO.CategoryID,
		Specification: specificationOutPut,
	}
}
