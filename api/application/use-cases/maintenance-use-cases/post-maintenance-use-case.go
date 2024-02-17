package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintUt "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/maintenance-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/rs/xid"
)

type PostMaintenanceUseCase struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewPostMaintenanceUseCase(
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) *PostMaintenanceUseCase {
	return &PostMaintenanceUseCase{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase *PostMaintenanceUseCase) ExecuteConcurrently(carID string, inputDTO m.MaintenanceInputDTO) (*m.MaintenanceOutputDTO, error) {

	// Criamos channels para o resultado e os erros
	resultChan := make(chan *m.MaintenanceOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool) // Um sinal para caso tenhamos um erro > true

	var wg sync.WaitGroup

	wg.Add(3)
	go useCase.performValidation(validationErrorSignal, &wg, errorChan, inputDTO)
	go useCase.checkCarStatus(validationErrorSignal, &wg, errorChan, carID, resultChan)
	go useCase.performMaintenanceCreation(validationErrorSignal, &wg, resultChan, errorChan, carID, inputDTO)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Esperamos pelos resultados ou erros
	select {
	case createdMaintenance := <-resultChan:
		return createdMaintenance, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *PostMaintenanceUseCase) performValidation(validationErrorSignal chan<- bool, wg *sync.WaitGroup, errorChan chan<- error, inputDTO m.MaintenanceInputDTO) {
	defer wg.Done()

	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
		errorChan <- err
		validationErrorSignal <- true
		return
	}

	validationErrorSignal <- false
}

func (useCase *PostMaintenanceUseCase) checkCarStatus(validationErrorSignal chan<- bool, wg *sync.WaitGroup, errorChan chan<- error, carID string, resultChan chan<- *m.MaintenanceOutputDTO) {
	defer wg.Done()

	car, err := useCase.carRepository.FindCarById(carID)
	if err != nil {
		errorChan <- err
		validationErrorSignal <- true
		return
	}

	if !car.Available {
		errorChan <- fmt.Errorf("car with id %s is not avaliable or in maintenance", carID)
		validationErrorSignal <- true
		return
	}

	validationErrorSignal <- false
}

func (useCase *PostMaintenanceUseCase) performMaintenanceCreation(validationErrorSignal chan<- bool, wg *sync.WaitGroup, resultChan chan<- *m.MaintenanceOutputDTO, errorChan chan<- error, carID string, inputDTO m.MaintenanceInputDTO) {
	defer wg.Done()

	newMaintenance, err := useCase.createMaintenanceInstance(carID, inputDTO)
	if err != nil {
		errorChan <- err
		validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
		return
	}

	go func() {
		if err := useCase.maintenanceRepository.CreateMaintenance(newMaintenance); err != nil {
			errorChan <- fmt.Errorf("failed to create maintenance record: %w", err)
			validationErrorSignal <- true // Sinaliza que ocorreu um erro de validação
			return
		}
	}()

	resultChan <- maintUt.ConvertToOutputDTO(newMaintenance)
	validationErrorSignal <- false

}

func (useCase *PostMaintenanceUseCase) createMaintenanceInstance(carID string, inputDTO m.MaintenanceInputDTO) (*domain.Maintenance, error) {
	newMaintenanceID := xid.New().String()

	parts := maintUt.ConvertPartsInputToDTO(inputDTO.Parts, newMaintenanceID)

	return &domain.Maintenance{
		ID:                        newMaintenanceID,
		CarID:                     carID,
		MaintenanceType:           inputDTO.MaintenanceType,
		OdometerReading:           inputDTO.OdometerReading,
		LastMaintenanceDate:       inputDTO.LastMaintenanceDate,
		ScheduledMaintenance:      inputDTO.ScheduledMaintenance,
		MaintenanceStatus:         inputDTO.MaintenanceStatus,
		MaintenanceDuration:       inputDTO.MaintenanceDuration,
		Description:               inputDTO.Description,
		MaintenanceNotes:          inputDTO.MaintenanceNotes,
		LaborCost:                 inputDTO.LaborCost,
		PartsCost:                 inputDTO.PartsCost,
		NextMaintenanceDueDate:    inputDTO.NextMaintenanceDueDate,
		MaintenanceCompletionDate: inputDTO.MaintenanceCompletionDate,
		Parts:                     parts,
	}, nil
}
