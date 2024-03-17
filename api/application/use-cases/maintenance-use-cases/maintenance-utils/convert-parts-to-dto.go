package maintenanceutils

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
	"github.com/rs/xid"
)

func ConvertPartsOutPutToDTO(parts []domain.Part) []m.PartOutputDTO {
	var partsDTO []m.PartOutputDTO
	for _, part := range parts {
		partsDTO = append(partsDTO, m.PartOutputDTO{
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
