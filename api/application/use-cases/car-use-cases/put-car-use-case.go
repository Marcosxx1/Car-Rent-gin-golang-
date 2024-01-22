package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
)

func PutCarUseCase(id string, registerRequest dtos.CarOutputDTO,
	carRepository repositories.CarRepository) (*dtos.CarOutputDTO, error) {

	carToBeUpdated := &domain.Car{
		Id:           registerRequest.Id,
		Name:         registerRequest.Name,
		Description:  registerRequest.Description,
		DailyRate:    registerRequest.DailyRate,
		Available:    registerRequest.Available,
		LicensePlate: registerRequest.LicensePlate,
		FineAmount:   registerRequest.FineAmount,
		Brand:        registerRequest.Brand,
		CategoryId:   registerRequest.CategoryId,
		CreatedAt:    registerRequest.CreatedAt,
	}

	if err := error_handling.ValidateStruct(carToBeUpdated); err != nil {
		return nil, err
	}

	car, err := carRepository.UpdateCar(id, *carToBeUpdated)
	if err != nil {
		return nil, err
	}

	carToBeReturned := &dtos.CarOutputDTO{
		Id:           car.Id,
		Name:         car.Name,
		Description:  car.Description,
		Available:    car.Available,
		LicensePlate: car.LicensePlate,
		FineAmount:   car.FineAmount,
		Brand:        car.Brand,
		CategoryId:   car.CategoryId,
	}

/* 	println(carToBeReturned)
	fmt.Printf("%+v\n", carToBeReturned) */

	return carToBeReturned, nil
}
