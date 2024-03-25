package carusecases

import (
	"errors"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils"
)

type GetCarByIdUseCase struct {
	carRepository           repositories.CarRepository
	specificationRepository repositories.SpecificationRepository
}

func NewFindCarByIdUseCase(
	carRepository repositories.CarRepository,
	specificationRepository repositories.SpecificationRepository) *GetCarByIdUseCase {

	return &GetCarByIdUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *GetCarByIdUseCase) Execute(id string) (*cardtos.CarOutputDTO, error) {
	existCar, err := useCase.carRepository.FindCarById(id)
	if err != nil {
		return nil, err
	}

	if existCar == nil {
		return nil, errors.New("car not found")
	}

	existSpecification, err := useCase.specificationRepository.FindAllSpecificationsByCarId(id)
	if err != nil {
		return nil, err
	}

	if existSpecification == nil {
		return nil, errors.New("specifications not found")
	}

	carToBeReturned := &cardtos.CarOutputDTO{
		ID:            existCar.ID,
		Name:          existCar.Name,
		Description:   existCar.Description,
		DailyRate:     existCar.DailyRate,
		Available:     existCar.Available,
		LicensePlate:  existCar.LicensePlate,
		FineAmount:    existCar.FineAmount,
		Brand:         existCar.Brand,
		Specification: utils.ConvertSpecificationToDTO(existSpecification),
	}

	return carToBeReturned, nil
}
