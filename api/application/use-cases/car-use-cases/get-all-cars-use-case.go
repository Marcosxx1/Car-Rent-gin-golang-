package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
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
		return nil, err
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
			Specification: repoutils.ConvertSpecificationToDTO(specifications),
		}

		outputDTO = append(outputDTO, dto)
	}
	return outputDTO, nil
}
