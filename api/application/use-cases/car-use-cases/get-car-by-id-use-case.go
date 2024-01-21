package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)
			
func GetCarByIdUseCase(id string, carRepository repositories.CarRepository) (*domain.Car, error){
	existCar, err := carRepository.FindCarById(id)
	if err != nil {
		return nil, err
	}
	return existCar, nil
}