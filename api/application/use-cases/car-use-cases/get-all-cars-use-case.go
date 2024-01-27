package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
)

func GetAllCarsUseCase(carRepository repositories.CarRepository, page, pageSize int) ([]*dtos.CarOutputDTO, error) {
	allCars, err := carRepository.FindAllCars(page, pageSize)
	if err != nil {
		return nil, err
	}

	outputDTO := make([]*dtos.CarOutputDTO, 0)

	for _, car := range allCars {
		dto := &dtos.CarOutputDTO{
			ID:           car.ID,
			Name:         car.Name,
			Description:  car.Description,
			DailyRate:    car.DailyRate,
			Available:    car.Available,
			LicensePlate: car.LicensePlate,
			FineAmount:   car.FineAmount,
			Brand:        car.Brand,
			CreatedAt:    car.CreatedAt,
		}

		outputDTO = append(outputDTO, dto)
	}
	return outputDTO, nil
}
