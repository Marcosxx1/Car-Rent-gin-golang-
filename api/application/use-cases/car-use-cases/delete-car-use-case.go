package usecases

import (
	"errors"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type DeleteCarUseCase struct {
	carRepository           r.CarRepository
	specificationRepository r.SpecificationRepository
}

func NewDeleteCarUseCase(
	carRepository r.CarRepository,
	specificationRepository r.SpecificationRepository) *DeleteCarUseCase {

	return &DeleteCarUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *DeleteCarUseCase) Execute(id string) error {
	existCar, err := useCase.carRepository.FindCarById(id)
	if err != nil {
		return err
	}

	if existCar == nil {
		return errors.New("car not found")
	}

	return useCase.carRepository.DeleteCar(id)
}
