package carusecases

import (
	"errors"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils"
)

type GetAllCarsUseCase struct {
	carRepository           repositories.CarRepository
	specificationRepository repositories.SpecificationRepository
}

func NewGetAllCarsUseCase(
	carRepository repositories.CarRepository,
	specificationRepository repositories.SpecificationRepository) *GetAllCarsUseCase {
	return &GetAllCarsUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *GetAllCarsUseCase) Execute(page, pageSize int) ([]*cardtos.CarOutputDTO, error) {
	allCars, err := useCase.carRepository.FindAllCars(page, pageSize)
	if err != nil {
		return nil, errors.New("error finding cars")
	}

	outputDTO := make([]*cardtos.CarOutputDTO, 0)

	for _, car := range allCars {
		specifications, err := useCase.specificationRepository.FindAllSpecificationsByCarId(car.ID)
		if err != nil {
			return nil, err
		}

		dto := &cardtos.CarOutputDTO{
			ID:            car.ID,
			Name:          car.Name,
			Description:   car.Description,
			DailyRate:     car.DailyRate,
			Available:     car.Available,
			LicensePlate:  car.LicensePlate,
			FineAmount:    car.FineAmount,
			Brand:         car.Brand,
			CategoryID:    car.CategoryID,
			Specification: utils.ConvertSpecificationToDTO(specifications),
		}

		outputDTO = append(outputDTO, dto)
	}
	return outputDTO, nil
}
