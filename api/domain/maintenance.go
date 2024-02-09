package domain

import (
	"time"

	"gorm.io/gorm"
)

type Maintenance struct {
	gorm.Model
	ID                        string    `json:"id"`
	CarID                     string    `json:"car_id"`
	MaintenanceType           string    `json:"maintenance_type" binding:"required"`
	OdometerReading           int       `json:"odometer_reading" binding:"required"`
	LastMaintenanceDate       time.Time `json:"last_maintenance_date" binding:"required"`
	ScheduledMaintenance      bool      `json:"scheduled_maintenance" binding:"required"`
	MaintenanceStatus         string    `json:"maintenance_status" binding:"required"`
	MaintenanceDuration       string    `json:"maintenance_duration" binding:"required"`
	Description               string    `json:"description" binding:"required"`
	MaintenanceNotes          string    `json:"maintenance_notes" binding:"required"`
	LaborCost                 int       `json:"labor_cost" binding:"required"`
	PartsCost                 int       `json:"parts_cost" binding:"required"`
	NextMaintenanceDueDate    time.Time `json:"next_maintenance_due_date" binding:"required"`
	MaintenanceCompletionDate time.Time `json:"maintenance_completion_date" binding:"required"`
	Car                       []*Car    `gorm:"many2many:car_maintenances"`
	Parts                     []Part    `gorm:"foreignkey:MaintenanceID"`
}

type Part struct {
	gorm.Model
	MaintenanceID   string    `json:"maintenance_id"`
	Name            string    `json:"name"`
	Cost            int       `json:"cost"`
	Quantity        int       `json:"quantity"`
	ReplacementDate time.Time `json:"replacement_date"`
}
