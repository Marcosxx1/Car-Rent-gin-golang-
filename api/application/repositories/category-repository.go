package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type CategoryPort interface {
	FindByName(name string) (*domain.Category, error)
    List(page, limit int) ([]*domain.Category, error)
    Create(category *domain.Category) (*domain.Category, error)
}