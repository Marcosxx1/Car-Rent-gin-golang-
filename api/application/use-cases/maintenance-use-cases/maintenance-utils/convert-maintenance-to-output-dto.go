package maintenanceutils

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
)

func ConvertMaintenanceListToDTOs(maintenanceList []*domain.Maintenance) []*m.MaintenanceOutputDTO {
	outputDTOList := make([]*m.MaintenanceOutputDTO, len(maintenanceList))

	for i, maintenance := range maintenanceList {
		outputDTOList[i] = ConvertToOutputDTO(maintenance)
	}

	return outputDTOList
}

func ConvertToOutputDTO(maintenance *domain.Maintenance) *m.MaintenanceOutputDTO {
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
		Parts:                     ConvertPartsOutPutToDTO(maintenance.Parts),
	}
}
