package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type CarRepository interface {
	RegisterCar(car *domain.Car) error
	FindCarByLicensePlate(licensePlate string) (*domain.Car, error)
	FindAllCars(page, pageSize int) ([]*domain.Car, error)
	DeleteCar(id string) error
	FindCarById(id string) (*domain.Car, error)
	UpdateCar(id string, car *domain.Car) (*domain.Car, error)
	AlterCarStatus(id string, available bool) error
	/*
		FindAvailableCars() (brand string, category_id string, name string, cars []*domain.Car)
		FindAvailableCarById(id string) *domain.Car
		UpdateAvailableCar(id string, available bool) error
	*/

}
