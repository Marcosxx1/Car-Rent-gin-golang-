package usecases

import (
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
		return err // implementar melhor erro
	}

	err = useCase.carRepository.DeleteCar(id)
	if err != nil {
		return err
	}
	return nil
}