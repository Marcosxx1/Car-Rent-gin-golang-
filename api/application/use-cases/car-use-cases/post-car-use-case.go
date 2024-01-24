package usecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
	"github.com/rs/xid"
)

func PostCarUseCase(
	registerRequest dtos.CarInputDTO,
	carRepository repositories.CarRepository) (*dtos.CarOutputDTO, error) {

	existCar, err := carRepository.FindCarByLicensePlate(registerRequest.LicensePlate)
	if err != nil {
		return nil, err
	}
	if existCar != nil {
		return nil, errors.New("car already exists")
	}

	newCar := &domain.Car{
		ID:           xid.New().String(),
		Name:         registerRequest.Name,
		Description:  registerRequest.Description,
		DailyRate:    registerRequest.DailyRate,
		Available:    registerRequest.Available,
		LicensePlate: registerRequest.LicensePlate,
		FineAmount:   registerRequest.FineAmount,
		Brand:        registerRequest.Brand,
	}

	if err := error_handling.ValidateStruct(newCar); err != nil {
		return nil, err
	}

	if err := carRepository.RegisterCar(newCar); err != nil {
		return nil, fmt.Errorf("failed to create car: %w", err)
	}

	outPut := &dtos.CarOutputDTO{
		ID:           newCar.ID,
		Name:         newCar.Name,
		Description:  newCar.Description,
		DailyRate:    newCar.DailyRate,
		Available:    newCar.Available,
		LicensePlate: newCar.LicensePlate,
		FineAmount:   newCar.FineAmount,
		Brand:        newCar.Brand,
	}

	return outPut, nil
}
