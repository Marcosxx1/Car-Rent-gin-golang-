package carusecases

import (
	"fmt"

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
	/* func (repositories.CarRepository) FindAllCars(page int, pageSize int) ([]*domain.Car, error) */
	allCars, err := useCase.carRepository.FindAllCars(page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("error finding cars for page %d, pageSize %d", page, pageSize)
	}

	if len(allCars) == 0 {
		return nil, fmt.Errorf("no cars found for page %d, pageSize %d", page, pageSize)
	}

	outputDTO := make([]*cardtos.CarOutputDTO, 0)

	for _, car := range allCars {
		/* func (repositories.SpecificationRepository) FindAllSpecificationsByCarId(carID string) ([]*domain.Specification, error) */
		specifications, err := useCase.specificationRepository.FindAllSpecificationsByCarId(car.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve specifications for car %s: %w", car.ID, err)
		}

		dto := &cardtos.CarOutputDTO{
			ID:           car.ID,
			Name:         car.Name,
			Description:  car.Description,
			DailyRate:    car.DailyRate,
			Available:    car.Available,
			LicensePlate: car.LicensePlate,
			FineAmount:   car.FineAmount,
			Brand:        car.Brand,
			CategoryID:   car.CategoryID,
			/* func (repositories.SpecificationRepository) FindAllSpecificationsByCarId(carID string) ([]*domain.Specification, error) */
			/*  */
			Specification: utils.ConvertSpecificationToDTO(specifications),
		}

		outputDTO = append(outputDTO, dto)
	}
	return outputDTO, nil
}
