package maintenanceusecases

import (
	"errors"
	"fmt"

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

func (useCase *PostMaintenanceUseCase) Execute(carID string, inputDTO m.MaintenanceInputDTO) (*m.MaintenanceOutputDTO, error) {

	newMaintenanceID := xid.New().String()
	var parts []domain.Part
	for _, partDTO := range inputDTO.Parts {
		parts = append(parts, domain.Part{
			Name:            partDTO.Name,
			Cost:            partDTO.Cost,
			Quantity:        partDTO.Quantity,
			ReplacementDate: partDTO.ReplacementDate,
		})
	}

	newMaintenance := &domain.Maintenance{
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
	}	
 
	if err := validation_errors.ValidateStruct(newMaintenance); err != nil {
		return nil, err
	}

	existingCar, err := useCase.carRepository.FindCarById(carID)
	if err != nil {
		return nil, err
	}
	if existingCar == nil {
		return nil, errors.New("car not found")
	}

	if err := useCase.maintenanceRepository.CreateMaintenance(newMaintenance); err != nil {
		return nil, fmt.Errorf("failed to create maintenance record: %w", err)
	}

	outputDTO := &m.MaintenanceOutputDTO{
		ID:                        newMaintenance.ID,
		CarID:                     newMaintenance.CarID,
		MaintenanceType:           newMaintenance.MaintenanceType,
		OdometerReading:           newMaintenance.OdometerReading,
		LastMaintenanceDate:       newMaintenance.LastMaintenanceDate,
		ScheduledMaintenance:      newMaintenance.ScheduledMaintenance,
		MaintenanceStatus:         newMaintenance.MaintenanceStatus,
		MaintenanceDuration:       newMaintenance.MaintenanceDuration,
		Description:               newMaintenance.Description,
		MaintenanceNotes:          newMaintenance.MaintenanceNotes,
		LaborCost:                 newMaintenance.LaborCost,
		PartsCost:                 newMaintenance.PartsCost,
		NextMaintenanceDueDate:    newMaintenance.NextMaintenanceDueDate,
		MaintenanceCompletionDate: newMaintenance.MaintenanceCompletionDate,
		Parts:                     convertPartsToDTO(newMaintenance.Parts),
	}

	return outputDTO, nil
}

func convertPartsToDTO(parts []domain.Part) []m.PartOutputDTO {
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
