package database

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGCarRepository struct{}

// RegisterCar registers a new car in the database.
// It takes a domain.Car as a parameter, creates a new record in the database,
// and returns a pointer to the registered car.
//
// Example:
//   newCar := domain.Car{
//       // set car properties
//   }
//   registeredCar := carRepository.RegisterCar(newCar)
//   fmt.Println("Registered Car:", registeredCar)
//
// Parameters:
//   - car: The car to be registered.
//
// Returns:
//   - *domain.Car: A pointer to the registered car.
func (repo *PGCarRepository) RegisterCar(car *domain.Car) error {
	return dbconfig.Postgres.Create(&car).Error
}


// FindCarByLicensePlate finds a car in the database based on its license plate.
// It takes a license plate as a parameter, queries the database, and returns
// a pointer to the found car along with an error.
//
// Example:
//   licensePlate := "ABC123"
//   foundCar, err := carRepository.FindCarByLicensePlate(licensePlate)
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else if foundCar == nil {
//       fmt.Println("Car not found.")
//   } else {
//       fmt.Println("Found Car:", foundCar)
//   }
//
// Parameters:
//   - licensePlate: The license plate of the car to be found.
//
// Returns:
//   - *domain.Car: A pointer to the found car, or nil if not found.
//   - error: An error, if any.
func (repo *PGCarRepository) FindCarByLicensePlate(licensePlate string) (*domain.Car, error) {
	var car domain.Car
	err := dbconfig.Postgres.Where("license_plate = ?", licensePlate).First(&car).Error

	if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, nil
			}
			return nil, err
	}

	return &car, nil
}


// FindAllCars retrieves all cars from the database.
// It returns a slice of pointers to cars and an error, if any.
//
// Example:
//   allCars, err := carRepository.FindAllCars()
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else {
//       fmt.Println("All Cars:", allCars)
//   }
//
// Returns:
//   - []*domain.Car: A slice of pointers to all cars in the database.
//   - error: An error, if any.


func (repo *PGCarRepository) FindAllCars(page, pageSize int) ([]*domain.Car, error) {
	var cars []*domain.Car
	offset := (page - 1) * pageSize

	err := dbconfig.Postgres.Offset(offset).Limit(pageSize).Find(&cars).Error

	if err != nil {
		return nil, err
	}

	return cars, nil
}


// DeleteCar deletes a car from the database based on its ID.
// It takes the ID of the car to be deleted, performs the deletion,
// and returns an error if the car is not found or if there's any issue.
//
// Example:
//   carID := "123"
//   err := carRepository.DeleteCar(carID)
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else {
//       fmt.Println("Car deleted successfully.")
//   }
//
// Parameters:
//   - id: The ID of the car to be deleted.
//
// Returns:
//   - error: An error, if any.
func (repo *PGCarRepository) DeleteCar(id string) error {
	result := dbconfig.Postgres.Where("id = ?", id).Delete(&domain.Car{})

	if result.Error != nil {
			return result.Error
	}

	if result.RowsAffected == 0 {
			return errors.New("car not found")
	}

	return nil
}


// FindCarById finds a car in the database based on its ID.
// It takes the ID as a parameter, queries the database, and returns
// a pointer to the found car along with an error.
//
// Example:
//   carID := "123"
//   foundCar, err := carRepository.FindCarById(carID)
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else if foundCar == nil {
//       fmt.Println("Car not found.")
//   } else {
//       fmt.Println("Found Car:", foundCar)
//   }
//
// Parameters:
//   - id: The ID of the car to be found.
//
// Returns:
//   - *domain.Car: A pointer to the found car, or nil if not found.
//   - error: An error, if any.
func (repo *PGCarRepository) FindCarById(id string) (*domain.Car, error) {
	var car domain.Car
	err := dbconfig.Postgres.Where("id = ?", id).First(&car).Error

	if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, nil
			}
			return nil, err
	}

	return &car, nil
}

// UpdateCar updates a car in the PostgreSQL database based on its ID.
// It takes the ID of the car to be updated and the updated car data.
// The function updates the corresponding record in the database and returns
// a pointer to the updated car along with an error.
//
// Example:
//   carID := "123"
//   updatedCarData := domain.Car{Name: "Updated Car", ...}
//   updatedCar, err := carRepository.UpdateCar(carID, updatedCarData)
//   if err != nil {
//       fmt.Println("Error:", err)
//   } else {
//       fmt.Println("Updated Car:", updatedCar)
//   }
//
// Parameters:
//   - id: The ID of the car to be updated.
//   - car: The updated car data.
//
// Returns:
//   - *domain.Car: A pointer to the updated car.
//   - error: An error, if any.
func (repo *PGCarRepository) UpdateCar(id string, car domain.Car) (*domain.Car, error) {
	dbconfig.Postgres.Model(&car).Where("id = ?", id).Updates(&car)
	
	return &car, nil
}
