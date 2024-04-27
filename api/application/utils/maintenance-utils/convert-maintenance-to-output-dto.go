package maintenanceutils

import (
	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

func ConvertMaintenanceListToDTOs(maintenanceList []*domain.Maintenance) []*maintenancedtos.MaintenanceOutputDTO {
	outputDTOList := make([]*maintenancedtos.MaintenanceOutputDTO, len(maintenanceList))

	for i, maintenance := range maintenanceList {
		outputDTOList[i] = ConvertToOutputDTO(maintenance)
	}

	return outputDTOList
}

func ConvertToOutputDTO(maintenance *domain.Maintenance) *maintenancedtos.MaintenanceOutputDTO {
	return &maintenancedtos.MaintenanceOutputDTO{
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
