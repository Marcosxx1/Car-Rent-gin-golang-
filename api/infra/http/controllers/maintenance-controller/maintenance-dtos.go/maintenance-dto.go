package maintenancedtos

import "time"

type MaintenanceInputDTO struct {
	CarID                     string         `json:"car_id"`
	MaintenanceType           string         `json:"maintenance_type" binding:"required"`
	OdometerReading           int            `json:"odometer_reading" binding:"required,min=0"`
	ScheduledMaintenance      bool           `json:"scheduled_maintenance"`
	MaintenanceStatus         string         `json:"maintenance_status" binding:"required"`
	MaintenanceDuration       string         `json:"maintenance_duration" binding:"omitempty"`
	Description               string         `json:"description"`
	MaintenanceNotes          string         `json:"maintenance_notes"`
	LaborCost                 int            `json:"labor_cost" binding:"min=1"`
	PartsCost                 int            `json:"parts_cost" binding:"min=1"`
	LastMaintenanceDate       time.Time      `json:"last_maintenance_date" form:"2006-01-02" binding:"required"`
	NextMaintenanceDueDate    time.Time      `json:"next_maintenance_due_date" form:"2006-01-02" binding:"required,gtfield=LastMaintenanceDate"`
	MaintenanceCompletionDate time.Time      `json:"maintenance_completion_date" form:"2006-01-02" binding:"required,gtfield=LastMaintenanceDate"`
	Parts                     []PartInputDTO `json:"parts"`
}

type MaintenanceOutputDTO struct {
	ID                        string    `json:"id"`
	CarID                     string    `json:"car_id"`
	MaintenanceType           string    `json:"maintenance_type"`
	OdometerReading           int       `json:"odometer_reading"`
	ScheduledMaintenance      bool      `json:"scheduled_maintenance"`
	MaintenanceStatus         string    `json:"maintenance_status"`
	MaintenanceDuration       string    `json:"maintenance_duration"`
	Description               string    `json:"description"`
	MaintenanceNotes          string    `json:"maintenance_notes"`
	LaborCost                 int       `json:"labor_cost"`
	PartsCost                 int       `json:"parts_cost"`
	LastMaintenanceDate       time.Time `json:"last_maintenance_date" form:"2006-01-02" binding:"required"`
	NextMaintenanceDueDate    time.Time `json:"next_maintenance_due_date" form:"2006-01-02" binding:"required"`
	MaintenanceCompletionDate time.Time `json:"maintenance_completion_date" form:"2006-01-02" binding:"required"`

	Parts []PartOutputDTO `json:"parts"`
}
 
type PartInputDTO struct {
	Name            string    `json:"name" binding:"required"`
	Cost            int       `json:"cost"`
	Quantity        int       `json:"quantity"`
	ReplacementDate time.Time `json:"replacement_date"`
}

type PartOutputDTO struct {
	ID              string    `json:"id"`
	MaintenanceID   string    `json:"maintenance_id"`
	Name            string    `json:"name"`
	Cost            int       `json:"cost"`
	Quantity        int       `json:"quantity"`
	ReplacementDate time.Time `json:"replacement_date"`
}
