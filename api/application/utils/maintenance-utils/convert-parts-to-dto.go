package maintenanceutils

import (
	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/rs/xid"
)

func ConvertPartsOutPutToDTO(parts []domain.Part) []maintenancedtos.PartOutputDTO {
	var partsDTO []maintenancedtos.PartOutputDTO
	for _, part := range parts {
		partsDTO = append(partsDTO, maintenancedtos.PartOutputDTO{
			ID:              part.ID,
			MaintenanceID:   part.MaintenanceID,
			Name:            part.Name,
			Cost:            part.Cost,
			Quantity:        part.Quantity,
			ReplacementDate: part.ReplacementDate,
		})
	}
	return partsDTO
}

func ConvertPartsInputToDTO(parts []maintenancedtos.PartInputDTO, newMaintenanceID string) []domain.Part {
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
