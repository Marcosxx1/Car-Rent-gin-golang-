package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type SpecificationRepository interface {
	FindSpecificationByName(name string) (*domain.Specification, error)
	GetAll() ([]*domain.Specification, error)
	PostSpecification(specification *domain.Specification) error
	FindAllSpecificationsByCarId(carID string) ([]*domain.Specification, error)
	UpdateSpecification(car_id string, specification []*domain.Specification) ([]*domain.Specification, error)
	PostMultipleSpecifications(specifications []*domain.Specification) error
}
