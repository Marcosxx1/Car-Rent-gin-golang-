package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type SpecificationRepository interface {
	FindSpecificationByName(name string) (*domain.Specification, error)
	GetAll() ([]*domain.Specification, error)
	PostSpecification(specification *domain.Specification) error
}
