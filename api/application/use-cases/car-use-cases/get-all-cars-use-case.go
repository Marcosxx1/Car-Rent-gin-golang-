package carusecases

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	utils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/utils"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
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

func (useCase *GetAllCarsUseCase) Execute(page, pageSize int) ([]*dtos.CarOutputDTO, error) {
	allCars, err := useCase.carRepository.FindAllCars(page, pageSize)
	if err != nil {
		return nil, errors.New("error finding cars")
	}

	outputDTO := make([]*dtos.CarOutputDTO, 0)

	for _, car := range allCars {
		specifications, err := useCase.specificationRepository.FindAllSpecificationsByCarId(car.ID)
		if err != nil {
			return nil, err
		}

		dto := &dtos.CarOutputDTO{
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
