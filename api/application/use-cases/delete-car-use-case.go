package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

func DeleteCarUseCase (carRepository repositories.CarRepository, id string) error {

	err := carRepository.DeleteCar(id)

	if err != nil {
		return err
	}
	return nil
}