package carusecases

import (
	"errors"

	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type DeleteCarUseCase struct {
	carRepository           repositories.CarRepository
	specificationRepository repositories.SpecificationRepository
}

func NewDeleteCarUseCase(
	carRepository repositories.CarRepository,
	specificationRepository repositories.SpecificationRepository) *DeleteCarUseCase {

	return &DeleteCarUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *DeleteCarUseCase) Execute(id string) error {

	existCar, err := useCase.carRepository.FindCarById(id)
	if err != nil {
		return errors.New("no car with provided id")
	}

	if existCar == nil {
		return errors.New("car not found")
	}

	return useCase.carRepository.DeleteCar(id)
}
