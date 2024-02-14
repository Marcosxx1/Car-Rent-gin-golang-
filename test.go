package maintenanceusecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
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
	newMaintenanceID := xid.New().String()
	fmt.Printf("+%v\n", carID)
	fmt.Printf("+%v\n", inputDTO)
	// Create channels for results and errors
	resultChan := make(chan *m.MaintenanceOutputDTO)
	errorChan := make(chan error)

	// Create WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Launch goroutines for each task
	wg.Add(2) // Number of tasks to be executed concurrently
	go useCase.performValidation(&wg, errorChan, newMaintenanceID, carID, inputDTO, resultChan)
	go useCase.performMaintenanceCreation(&wg, resultChan, errorChan, newMaintenanceID, carID, inputDTO)

	// Start a goroutine to close channels and wait for all tasks to finish
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Wait for results or errors
	select {
	case createdMaintenance := <-resultChan:
		return createdMaintenance, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *PostMaintenanceUseCase) performValidation(wg *sync.WaitGroup, errorChan chan<- error, newMaintenanceID, carID string, inputDTO m.MaintenanceInputDTO, resultChan chan<- *m.MaintenanceOutputDTO) {
	defer wg.Done()

	newMaintenance, err := useCase.createMaintenanceInstance(newMaintenanceID, carID, inputDTO)
	if err != nil {
		errorChan <- err
		return
	}

	// Validate maintenance using a separate goroutine
	go func() {//        C:/Users/marcos.s/Documents/prjetos/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/post-maintenance-use-case.go:70 +0x165
		if err := validation_errors.ValidateStruct(newMaintenance); err != nil {
			errorChan <- err
			return
		}
		resultChan <- convertToOutputDTO(newMaintenance) //        C:/Users/marcos.s/Documents/prjetos/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases/post-maintenance-use-case.go:75 +0x79
	}()
}

func (useCase *PostMaintenanceUseCase) performMaintenanceCreation(wg *sync.WaitGroup, resultChan chan<- *m.MaintenanceOutputDTO, errorChan chan<- error, newMaintenanceID, carID string, inputDTO m.MaintenanceInputDTO) {
	defer wg.Done()

	newMaintenance, err := useCase.createMaintenanceInstance(newMaintenanceID, carID, inputDTO)
	if err != nil {
		errorChan <- err
		return
	}

	// Create maintenance using a separate goroutine
	go func() {
		if err := useCase.maintenanceRepository.CreateMaintenance(newMaintenance); err != nil {
			errorChan <- fmt.Errorf("failed to create maintenance record: %w", err)
			return
		}
		resultChan <- convertToOutputDTO(newMaintenance)
	}()
}

func (useCase *PostMaintenanceUseCase) createMaintenanceInstance(newMaintenanceID, carID string, inputDTO m.MaintenanceInputDTO) (*domain.Maintenance, error) {
	parts := ConvertPartsInputToDTO(inputDTO.Parts)

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

func convertToOutputDTO(maintenance *domain.Maintenance) *m.MaintenanceOutputDTO {
	return &m.MaintenanceOutputDTO{
		ID:                        maintenance.ID,
		CarID:                     maintenance.CarID,
		MaintenanceType:           maintenance.MaintenanceType,
		OdometerReading:           maintenance.OdometerReading,
		LastMaintenanceDate:       maintenance.LastMaintenanceDate,
		ScheduledMaintenance:      maintenance.ScheduledMaintenance,
		MaintenanceStatus:         maintenance.MaintenanceStatus,
		MaintenanceDuration:       maintenance.MaintenanceDuration,
		Description:               maintenance.Description,
		MaintenanceNotes:          maintenance.MaintenanceNotes,
		LaborCost:                 maintenance.LaborCost,
		PartsCost:                 maintenance.PartsCost,
		NextMaintenanceDueDate:    maintenance.NextMaintenanceDueDate,
		MaintenanceCompletionDate: maintenance.MaintenanceCompletionDate,
		Parts:                     convertPartsOutPutToDTO(maintenance.Parts),
	}
}
func convertPartsOutPutToDTO(parts []domain.Part) []m.PartOutputDTO {
	var partsDTO []m.PartOutputDTO
	for _, part := range parts {
		partsDTO = append(partsDTO, m.PartOutputDTO{
			MaintenanceID:   part.MaintenanceID,
			Name:            part.Name,
			Cost:            part.Cost,
			Quantity:        part.Quantity,
			ReplacementDate: part.ReplacementDate,
		})
	}
	return partsDTO
}

func ConvertPartsInputToDTO(parts []m.PartInputDTO) []domain.Part {
	var partsDTO []domain.Part
	for _, part := range parts {
		partsDTO = append(partsDTO, domain.Part{
			Name:            part.Name,
			Cost:            part.Cost,
			Quantity:        part.Quantity,
			ReplacementDate: part.ReplacementDate,
		})
	}
	return partsDTO
}
