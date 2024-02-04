package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type MaintenanceRepository interface {
	// CreateMaintenance adiciona uma nova manutenção ao repositório
	CreateMaintenance(maintenance *domain.Maintenance) error

	// GetMaintenanceByID retorna uma manutenção com base no ID
	GetMaintenanceByID(id string) (*domain.Maintenance, error)

	// UpdateMaintenance atualiza as informações de uma manutenção existente
	UpdateMaintenance(maintenance *domain.Maintenance) error

	// DeleteMaintenance remove uma manutenção do repositório com base no ID
	DeleteMaintenance(id string) error

	// GetMaintenancesByCarID retorna todas as manutenções associadas a um carro específico
	GetMaintenancesByCarID(carID string) ([]*domain.Maintenance, error)

	// GetMaintenancesByType retorna todas as manutenções de um determinado tipo
	GetMaintenancesByType(maintenanceType string) ([]*domain.Maintenance, error)

	// GetScheduledMaintenances retorna todas as manutenções programadas
	GetScheduledMaintenances() ([]*domain.Maintenance, error)

	// GetCompletedMaintenances retorna todas as manutenções concluídas
	GetCompletedMaintenances() ([]*domain.Maintenance, error)

	// GetMaintenanceByStatus retorna todas as manutenções com um determinado status
	GetMaintenancesByStatus(status string) ([]*domain.Maintenance, error)

	// GetMaintenancesByDateRange retorna todas as manutenções dentro de um intervalo de datas
	GetMaintenancesByDateRange(startDate, endDate string) ([]*domain.Maintenance, error)
}
