package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

func GetAllCarsUseCase(carRepository repositories.CarRepository) ([]*domain.Car, error){

	allCars, err := carRepository.FindAllCars()

	if err != nil {
		return nil, err
	}

	return allCars, nil

}