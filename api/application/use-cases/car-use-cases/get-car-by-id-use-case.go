package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
)

func GetCarByIdUseCase(id string, carRepository repositories.CarRepository) (*dtos.CarOutputDTO, error) {
	existCar, err := carRepository.FindCarById(id)
	if err != nil {
		return nil, err
	}

	carToBeReturned := &dtos.CarOutputDTO{
		ID:           existCar.ID,
		Name:         existCar.Name,
		Description:  existCar.Description,
		DailyRate:    existCar.DailyRate,
		Available:    existCar.Available,
		LicensePlate: existCar.LicensePlate,
		FineAmount:   existCar.FineAmount,
		Brand:        existCar.Brand,
		CreatedAt:    existCar.CreatedAt,
	}

	return carToBeReturned, nil
}
