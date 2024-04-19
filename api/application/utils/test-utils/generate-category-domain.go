package testutils

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

func GenerateDomainCategories(numCategories int) []*domain.Category {

	var categories []*domain.Category

	for i := 0; i < numCategories; i++ {
		id := fmt.Sprintf("id%d", i)
		name := fmt.Sprintf("name%d", i)
		description := fmt.Sprintf("description for category id%d,", i)
		carID := fmt.Sprintf("carID%d,", i)

		car := &domain.Car{
			ID: carID,
		}

		category := &domain.Category{
			ID:          id,
			Name:        name,
			Description: description,
			Car:         &[]domain.Car{*car},
		}

		categories = append(categories, category)

	}
	return categories
}
