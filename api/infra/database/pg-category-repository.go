package database

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGCategory struct{}

func NewPGCategoryRepository() repositories.CategoryRepository {
	return &PGCategory{}
}

// PostCategory creates a new category record in the database.
// It takes a domain.Category as a parameter, creates a new record in the database,
// and returns an error.
//
// Example:
//
//	newCategory := domain.Category{
//	    // set category properties
//	}
//	err := categoryRepository.PostCategory(newCategory)
//	if err != nil {
//	    // handle error
//	}
//
// Parameters:
//   - category: The category to be created.
//
// Returns:
//   - error: An error, if any, during the creation process.
func (repo *PGCategory) PostCategory(category *domain.Category) error {
	return dbconfig.Postgres.Create(category).Error
}

// FindCategoryByName retrieves a category from the database by its name.
// It takes a string (name) as a parameter, queries the database for a category with that name,
// and returns a pointer to the category or nil if the category is not found. It also returns an error.
//
// Example:
//
//	categoryName := "example_category"
//	foundCategory, err := categoryRepository.FindCategoryByName(categoryName)
//	if err != nil {
//	    // handle error
//	}
//	if foundCategory != nil {
//	    fmt.Println("Found Category:", *foundCategory)
//	} else {
//	    fmt.Println("Category not found.")
//	}
//
// Parameters:
//   - name: The name of the category to be retrieved.
//
// Returns:
//   - *domain.Category: A pointer to the retrieved category, or nil if not found.
//   - error: An error, if any, during the retrieval process.
func (repo *PGCategory) FindCategoryByName(name string) (*domain.Category, error) {
	var category domain.Category
	err := dbconfig.Postgres.Where("name = ?", name).First(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Category not found, return nil
		}
		return nil, err // Other errors
	}

	return &category, nil
}

// GetAll retrieves all categories from the database.
// It returns a slice of pointers to domain.Category and an error.
//
// Example:
//
//	allCategories, err := categoryRepository.GetAll()
//	if err != nil {
//	    // handle error
//	}
//	for _, cat := range allCategories {
//	    fmt.Println("Category:", *cat)
//	}
//
// Returns:
//   - []*domain.Category: A slice of pointers to all categories in the database.
//   - error: An error, if any, during the retrieval process.
func (repo *PGCategory) GetAll() ([]*domain.Category, error) {
	var categories []*domain.Category
	err := dbconfig.Postgres.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}
