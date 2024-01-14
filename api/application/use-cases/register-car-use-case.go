package usecases

import (
	"time"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain/error_handling"
	"github.com/rs/xid"
)

type RegisterCarRequest struct {
	Name         string  
	Description  string  
	DailyRate    float64 
	Available    bool    
	LicensePlate string  
	FineAmount   float64 
	Brand        string  
	CategoryId   string  
}

func RegisterCarUseCase(
	registerRequest RegisterCarRequest,
	carRepository repositories.CarRepository) (*domain.Car, error) {

	car := &domain.Car{
			Id:           xid.New().String(),
			Name:         registerRequest.Name,
			Description:  registerRequest.Description,
			DailyRate:    registerRequest.DailyRate,
			Available:    registerRequest.Available,
			LicensePlate: registerRequest.LicensePlate,
			FineAmount:   registerRequest.FineAmount,
			Brand:        registerRequest.Brand,
			CategoryId:   registerRequest.CategoryId,
			CreatedAt:    time.Now(),
	}
 
	if err := error_handling.ValidateStruct(car); err != nil {
			return nil, err
	}

	return carRepository.RegisterCar(*car), nil
}
