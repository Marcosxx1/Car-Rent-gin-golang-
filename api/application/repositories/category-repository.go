package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type CategoryRepository interface {
	FindCategoryByName(name string) (*domain.Category, error)
   // GetAll(page, limit int) ([]*domain.Category, error)
    PostCategory(category domain.Category) (*domain.Category, error)
}