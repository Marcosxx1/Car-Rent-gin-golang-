package domain

import (
	"time"

	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Maintenance struct {
	gorm.Model
	ID                        string    `json:"id"`
	CarID                     string    `json:"car_id"`
	MaintenanceType           string    `json:"maintenance_type" binding:"required"`
	OdometerReading           int       `json:"odometer_reading" binding:"required"`
	LastMaintenanceDate       time.Time `json:"last_maintenance_date" binding:"required" validate:"gte=0"`
	ScheduledMaintenance      bool      `json:"scheduled_maintenance" binding:"required"`
	MaintenanceStatus         string    `json:"maintenance_status" binding:"required"`
	MaintenanceDuration       string    `json:"maintenance_duration" binding:"required"`
	Description               string    `json:"description" binding:"required"`
	MaintenanceNotes          string    `json:"maintenance_notes" binding:"required"`
	LaborCost                 int       `json:"labor_cost" binding:"required" validate:"gte=0"`
	PartsCost                 int       `json:"parts_cost" binding:"required" validate:"gte=0"`
	NextMaintenanceDueDate    time.Time `json:"next_maintenance_due_date" binding:"required"`
	MaintenanceCompletionDate time.Time `json:"maintenance_completion_date" binding:"required"`
	Car                       []*Car    `gorm:"many2many:car_maintenances"`
	Parts                     []Part    `gorm:"foreignkey:MaintenanceID"`
}

type Part struct {
	gorm.Model
	ID              string    `json:"id"`
	MaintenanceID   string    `json:"maintenance_id"`
	Name            string    `json:"name"`
	Cost            int       `json:"cost" validate:"gte=0"`
	Quantity        int       `json:"quantity" validate:"gte=0"`
	ReplacementDate time.Time `json:"replacement_date"`
}

func CreateMaintenanceInstance(carID string, inputDTO m.MaintenanceInputDTO) (*Maintenance, error) {
	newMaintenanceID := xid.New().String()

	parts := ConvertPartsInputToDTO(inputDTO.Parts, newMaintenanceID)

	return &Maintenance{
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

func ConvertPartsOutPutToDTO(parts []Part) []m.PartOutputDTO {
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

func ConvertPartsInputToDTO(parts []m.PartInputDTO, newMaintenanceID string) []Part {
	var partsDTO []Part
	for _, part := range parts {
		partsDTO = append(partsDTO, Part{
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
