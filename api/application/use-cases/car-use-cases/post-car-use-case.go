package usecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/rs/xid"
)

func PostCarUseCase(
	registerRequest dtos.CarInputDTO,
	carRepository repositories.CarRepository) (*dtos.CarOutputDTO, error) {

	var specifications []domain.Specification
	for _, specification := range registerRequest.Specification {
		specifications = append(specifications, domain.Specification{
			ID:          xid.New().String(),
			Name:        specification.Name,
			Description: specification.Description,
		})
	}

	newCar := &domain.Car{
		ID:            xid.New().String(),
		Name:          registerRequest.Name,
		Description:   registerRequest.Description,
		DailyRate:     registerRequest.DailyRate,
		Available:     registerRequest.Available,
		LicensePlate:  registerRequest.LicensePlate,
		FineAmount:    registerRequest.FineAmount,
		Brand:         registerRequest.Brand,
		CategoryID:    registerRequest.CategoryID,
		Specification: specifications,
	}

	if err := validation_errors.ValidateStruct(newCar); err != nil {
		return nil, err
	}

	existCar, err := carRepository.FindCarByLicensePlate(registerRequest.LicensePlate)
	if err != nil {
		return nil, err
	}
	if existCar != nil {
		return nil, errors.New("car already exists")
	}

	if err := carRepository.RegisterCar(newCar); err != nil {
		return nil, fmt.Errorf("A car needs a category to be registrated")
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
		CategoryID:   newCar.CategoryID,
	}

	return outPut, nil
}
