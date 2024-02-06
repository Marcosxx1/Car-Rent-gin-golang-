package domain

import (
	"time"

	"gorm.io/gorm"
)

type CarMaintenance struct {
	gorm.Model
	CarID         string `gorm:"primaryKey"`
	MaintenanceID string `gorm:"primaryKey"`
	ServiceDate   time.Time
	Description   string
}
