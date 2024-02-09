package usecases

import (
	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
)

type GetCarByIdUseCase struct {
	carRepository           r.CarRepository
	specificationRepository r.SpecificationRepository
}

func NewFindCarByIdUseCase(
	carRepository r.CarRepository,
	specificationRepository r.SpecificationRepository) *GetCarByIdUseCase {

	return &GetCarByIdUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *GetCarByIdUseCase) Execute(id string) (*dtos.CarOutputDTO, error) {
	existCar, err := useCase.carRepository.FindCarById(id)
	if err != nil {
		return nil, err
	}

	existSpecification, err := useCase.specificationRepository.FindAllSpecificationsByCarId(id)
	if err != nil {
		return nil, nil
	}

	carToBeReturned := &dtos.CarOutputDTO{
		ID:            existCar.ID,
		Name:          existCar.Name,
		Description:   existCar.Description,
		DailyRate:     existCar.DailyRate,
		Available:     existCar.Available,
		LicensePlate:  existCar.LicensePlate,
		FineAmount:    existCar.FineAmount,
		Brand:         existCar.Brand,
		Specification: repoutils.ConvertSpecificationToDTO(existSpecification),
	}

	return carToBeReturned, nil
}

