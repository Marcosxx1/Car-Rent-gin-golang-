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
	OdometerReading           int       `json:"odometer_reading"`
	LastMaintenanceDate       time.Time `json:"last_maintenance_date"`
	ScheduledMaintenance      bool      `json:"scheduled_maintenance"`
	MaintenanceStatus         string    `json:"maintenance_status" binding:"required"`
	MaintenanceDuration       string    `json:"maintenance_duration"`
	Description               string    `json:"description"`
	MaintenanceNotes          string    `json:"maintenance_notes"`
	LaborCost                 int       `json:"labor_cost"`
	PartsCost                 int       `json:"parts_cost"`
	NextMaintenanceDueDate    time.Time `json:"next_maintenance_due_date"`
	MaintenanceCompletionDate time.Time `json:"maintenance_completion_date"`
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
