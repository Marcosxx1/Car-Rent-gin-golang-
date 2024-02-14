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

	// Criamos channels para o resultado e os erros
	resultChan := make(chan *m.MaintenanceOutputDTO)
	errorChan := make(chan error)

	// Criamos um waitgroup para aguardar a finalização das goroutines
	var wg sync.WaitGroup

	// Criamos um Once para garantir que o fechamento dos canais ocorra apenas uma vez
	var once sync.Once
	closeChannels := func() {
		once.Do(func() {
			close(resultChan)
			close(errorChan)
		})
	}

	// Criamos goroutines para cada tarefa
	wg.Add(3)
	go useCase.performValidation(&wg, errorChan, newMaintenanceID, carID, inputDTO, resultChan)
	go useCase.performMaintenanceCreation(&wg, resultChan, errorChan, newMaintenanceID, carID, inputDTO)

	// Iniciamos a goroutine para fechar os canais e aguardar todas as tarefas terminarem
	go func() {
		wg.Wait()
		closeChannels()
	}()

	// Esperamos pelos resultados ou erros
	select {
	case createdMaintenance := <-resultChan:
		return createdMaintenance, nil
	case err := <-errorChan:
		closeChannels() // Fechamos os canais caso ocorra um erro antes de esperar a conclusão
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

	go func() {
		if err := validation_errors.ValidateStruct(newMaintenance); err != nil {
			errorChan <- err
			return
		}
		resultChan <- convertToOutputDTO(newMaintenance)

	}()
}

func (useCase *PostMaintenanceUseCase) performMaintenanceCreation(wg *sync.WaitGroup, resultChan chan<- *m.MaintenanceOutputDTO, errorChan chan<- error, newMaintenanceID, carID string, inputDTO m.MaintenanceInputDTO) {
	defer wg.Done()

	newMaintenance, err := useCase.createMaintenanceInstance(newMaintenanceID, carID, inputDTO)

	if err != nil {
		errorChan <- err
		return
	}

	go func() {
		if err := useCase.maintenanceRepository.CreateMaintenance(newMaintenance); err != nil {
			errorChan <- fmt.Errorf("failed to create maintenance record: %w", err)
			return
		}
		resultChan <- convertToOutputDTO(newMaintenance)

	}()
}

func (useCase *PostMaintenanceUseCase) createMaintenanceInstance(newMaintenanceID, carID string, inputDTO m.MaintenanceInputDTO) (*domain.Maintenance, error) {
	parts := ConvertPartsInputToDTO(inputDTO.Parts, newMaintenanceID)

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

func ConvertPartsInputToDTO(parts []m.PartInputDTO, newMaintenanceID string) []domain.Part {
	var partsDTO []domain.Part
	for _, part := range parts {
		partsDTO = append(partsDTO, domain.Part{
			ID:              xid.New().String(),
			MaintenanceID:   newMaintenanceID,
			Name:            part.Name,
			Cost:            part.Cost,
			Quantity:        part.Quantity,
			ReplacementDate: part.ReplacementDate,
		})
	}
	return partsDTO
}
