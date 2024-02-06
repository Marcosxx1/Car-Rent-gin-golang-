package maintenancedtosgo

import (
	"time"
)

type MaintenanceInputDTO struct {
	CarID                     string         `json:"car_id"`
	MaintenanceType           string         `json:"maintenance_type" binding:"required"`
	OdometerReading           int            `json:"odometer_reading"`
	LastMaintenanceDate       time.Time      `json:"last_maintenance_date"`
	ScheduledMaintenance      bool           `json:"scheduled_maintenance"`
	MaintenanceStatus         string         `json:"maintenance_status" binding:"required"`
	MaintenanceDuration       string         `json:"maintenance_duration"`
	Description               string         `json:"description"`
	PartsReplaced             []PartInputDTO `json:"parts_replaced"`
	MaintenanceNotes          string         `json:"maintenance_notes"`
	LaborCost                 int            `json:"labor_cost"`
	PartsCost                 int            `json:"parts_cost"`
	NextMaintenanceDueDate    time.Time      `json:"next_maintenance_due_date"`
	MaintenanceCompletionDate time.Time      `json:"maintenance_completion_date"`
}

type PartInputDTO struct {
	MaintenanceID string `json:"id"`
	Name          string `json:"name"`
	Quantity      int    `json:"quantity"`
	Cost          int    `json:"cost"`
}

type MaintenanceOutputDTO struct {
	ID                        string          `json:"id"`
	CarID                     string          `json:"car_id"`
	MaintenanceType           string          `json:"maintenance_type"`
	OdometerReading           int             `json:"odometer_reading"`
	LastMaintenanceDate       time.Time       `json:"last_maintenance_date"`
	ScheduledMaintenance      bool            `json:"scheduled_maintenance"`
	MaintenanceStatus         string          `json:"maintenance_status"`
	MaintenanceDuration       string          `json:"maintenance_duration"`
	Description               string          `json:"description"`
	PartsReplaced             []PartOutputDTO `json:"parts_replaced"`
	MaintenanceNotes          string          `json:"maintenance_notes"`
	LaborCost                 int             `json:"labor_cost"`
	PartsCost                 int             `json:"parts_cost"`
	NextMaintenanceDueDate    time.Time       `json:"next_maintenance_due_date"`
	MaintenanceCompletionDate time.Time       `json:"maintenance_completion_date"`
}

type PartOutputDTO struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Cost     int    `json:"cost"`
}
