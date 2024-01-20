package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
)

func PutCarUseCase(id string, car *domain.Car,
	 carRepository repositories.CarRepository)(*domain.Car, error){

		if err := error_handling.ValidateStruct(car); err != nil {
			return nil, err
	}

	car, err := carRepository.UpdateCar(id, *car)
	if err != nil {
		return nil, err
	}
	return car, nil
}