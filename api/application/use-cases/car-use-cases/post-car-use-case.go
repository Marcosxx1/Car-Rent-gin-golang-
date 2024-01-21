package usecases

import (
	"errors"
	"time"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
	"github.com/rs/xid"
)

func PostCarUseCase(
	registerRequest dtos.CarDto,
	carRepository repositories.CarRepository) (*domain.Car, error) {

		existCar, err := carRepository.FindCarByLicensePlate(registerRequest.LicensePlate)
		if err != nil {
				return nil, err
		}
		if existCar != nil {
				return nil, errors.New("car already exists")
		}

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
