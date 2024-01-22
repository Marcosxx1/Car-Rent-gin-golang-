package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
)

func GetAllCarsUseCase(carRepository repositories.CarRepository) ([]*dtos.CarOutputDTO, error){

 

	allCars, err := carRepository.FindAllCars()
	if err != nil {
			// Handle the error, return or log it
			return nil, err
	}
	
	outputDTO := make([]*dtos.CarOutputDTO, 0)
	for _, car := range allCars {
			dto := &dtos.CarOutputDTO{
					Id:           car.Id,  
					Name:         car.Name,
					Description:  car.Description,
					DailyRate:    car.DailyRate,
					Available:    car.Available,
					LicensePlate: car.LicensePlate,
					FineAmount:   car.FineAmount,
					Brand:        car.Brand,
					CategoryId:   car.CategoryId,
					CreatedAt:    car.CreatedAt,
			}
	
			outputDTO = append(outputDTO, dto)
	}
	return outputDTO, nil

}