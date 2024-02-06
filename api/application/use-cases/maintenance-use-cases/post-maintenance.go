package maintenanceusecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type PostMaintenanceUseCase struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewPostMaintenanceUseCase(
	context *gin.Context,
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) *PostMaintenanceUseCase {
	return &PostMaintenanceUseCase{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (repo *PostMaintenanceUseCase) Execute(carID string, inputDTO m.MaintenanceInputDTO) (*m.MaintenanceOutputDTO, error) {
	existCar, err := repo.carRepository.FindCarById(carID)
	if err != nil {
		return nil, err
	}
	if existCar == nil {
		return nil, errors.New("car not found")
	}

	maintenanceId := xid.New().String()

	var parts []domain.Part

	for _, part := range inputDTO.PartsReplaced {
		parts = append(parts, domain.Part{
			MaintenanceID: maintenanceId,
			Name:          part.Name,
			Quantity:      part.Quantity,
			Cost:          part.Cost,
		})
	}

	newMaintenance := &domain.Maintenance{
		ID:                        maintenanceId,
		CarID:                     carID,
		MaintenanceType:           inputDTO.MaintenanceType,
		OdometerReading:           inputDTO.OdometerReading,
		LastMaintenanceDate:       inputDTO.LastMaintenanceDate,
		ScheduledMaintenance:      inputDTO.ScheduledMaintenance,
		MaintenanceStatus:         inputDTO.MaintenanceStatus,
		MaintenanceDuration:       inputDTO.MaintenanceDuration,
		Description:               inputDTO.Description,
		PartsReplaced:             parts,
		MaintenanceNotes:          inputDTO.MaintenanceNotes,
		LaborCost:                 inputDTO.LaborCost,
		PartsCost:                 inputDTO.PartsCost,
		NextMaintenanceDueDate:    inputDTO.NextMaintenanceDueDate,
		MaintenanceCompletionDate: inputDTO.MaintenanceCompletionDate,
	}

	if err := error_handling.ValidateStruct(newMaintenance); err != nil {
		return nil, err
	}

	if err := repo.maintenanceRepository.CreateMaintenance(newMaintenance); err != nil {
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
		PartsReplaced:             []m.PartOutputDTO{},
		MaintenanceNotes:          newMaintenance.MaintenanceNotes,
		LaborCost:                 newMaintenance.LaborCost,
		PartsCost:                 newMaintenance.PartsCost,
		NextMaintenanceDueDate:    newMaintenance.NextMaintenanceDueDate,
		MaintenanceCompletionDate: newMaintenance.MaintenanceCompletionDate,
	}

	return outputDTO, nil
}
