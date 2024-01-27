package dtos

import "time"

type NewCarRequest struct {
	Name         string  
	Description  string  
	DailyRate    float64 
	Available    bool    
	LicensePlate string  
	FineAmount   float64 
	Brand        string  
	CategoryID   string  
	CreatedAt    time.Time 
 }