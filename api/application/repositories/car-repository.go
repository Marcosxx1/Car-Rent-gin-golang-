package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type CarRepository interface {
	RegisterCar(car domain.Car) *domain.Car
	FindCarByLicensePlate(licensePlate string) (*domain.Car, error)
	FindAllCars() ([]*domain.Car, error)
	DeleteCar(id string) error

	/*
		UpdateCar(id string, car domain.Car) *domain.Car
		FindCarById(id string) *domain.Car
		FindAvailableCars() (brand string, category_id string, name string, cars []*domain.Car)
		FindAvailableCarById(id string) *domain.Car
		UpdateAvailableCar(id string, available bool) error
	*/

}
