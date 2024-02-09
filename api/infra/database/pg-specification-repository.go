package database

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGSpecification struct{}

// FindAllSpecificationsById retrieves all specifications for a given car_id from the database.
// It takes a string (car_id) as a parameter, queries the database for specifications with that car_id,
// and returns a slice of pointers to the specifications or an error.
//
// Example:
//   carID := "example_car_id"
//   specifications, err := specificationRepository.FindAllSpecificationsById(carID)
//   if err != nil {
//       // handle error
//   }
//   for _, spec := range specifications {
//       fmt.Println("Specification:", *spec)
//   }
//
// Parameters:
//   - carID: The ID of the car for which specifications are to be retrieved.
//
// Returns:
//   - []*domain.Specification: A slice of pointers to specifications for the given car_id.
//   - error: An error, if any, during the retrieval process.
func (repo *PGSpecification) FindAllSpecificationsByCarId(carID string) ([]*domain.Specification, error) {
	var specifications []*domain.Specification

	err := dbconfig.Postgres.Where("car_id = ?", carID).Find(&specifications).Error

	if err != nil {
		return nil, err
	}

	return specifications, nil
}


// FindSpecificationByName retrieves a specification from the database by its name.
// It takes a string (name) as a parameter, queries the database for a specification with that name,
// and returns a pointer to the specification or nil if the specification is not found. It also returns an error.
//
// Example:
//   specificationName := "example_specification"
//   foundSpecification, err := specificationRepository.FindSpecificationByName(specificationName)
//   if err != nil {
//       // handle error
//   }
//   if foundSpecification != nil {
//       fmt.Println("Found Specification:", *foundSpecification)
//   } else {
//       fmt.Println("Specification not found.")
//   }
//
// Parameters:
//   - name: The name of the specification to be retrieved.
//
// Returns:
//   - *domain.Specification: A pointer to the retrieved specification, or nil if not found.
//   - error: An error, if any, during the retrieval process.
func (repo *PGSpecification) FindSpecificationByName(name string) (*domain.Specification, error) {
	var specification domain.Specification
	err := dbconfig.Postgres.Where("name = ?", name).First(&specification).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Specification not found, return nil
	} else if err != nil {
		return nil, err // Other errors
	}
	return &specification, nil
}



// GetAll retrieves all specifications from the database.
// It returns a slice of pointers to domain.Specification and an error.
//
// Example:
//   allSpecifications, err := specificationRepository.GetAll()
//   if err != nil {
//       // handle error
//   }
//   for _, spec := range allSpecifications {
//       fmt.Println("Specification:", *spec)
//   }
//
// Returns:
//   - []*domain.Specification: A slice of pointers to all specifications in the database.
//   - error: An error, if any, during the retrieval process.
func (repo *PGSpecification) GetAll() ([]*domain.Specification, error) {
	var specifications []*domain.Specification
	err := dbconfig.Postgres.Find(&specifications).Error

	if err != nil {
		return nil, err
	}
	return specifications, nil
}

// PostSpecification creates a new specification record in the database.
// It takes a domain.Specification as a parameter, creates a new record in the database,
// and returns an error.
//
// Example:
//   newSpecification := domain.Specification{
//       // set specification properties
//   }
//   err := specificationRepository.PostSpecification(newSpecification)
//   if err != nil {
//       // handle error
//   }
//
// Parameters:
//   - specification: The specification to be created.
//
// Returns:
//   - error: An error, if any, during the creation process.
func (repo *PGSpecification) PostSpecification(specification *domain.Specification) error {
	return dbconfig.Postgres.Create(specification).Error
}
