package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/category-controller/category-dtos"
)

func PostCategoryUseCase(
	registerCategory dtos.CategoryInputDTO,
	categoryRepository repositories.CategoryRepository)(*domain.Category, error) {

 

		category := &domain.Category{
			Name: registerCategory.Name,
			Description: registerCategory.Description,
		}

		if err := error_handling.ValidateStruct(category); err != nil {
			return nil, err
	}

	return categoryRepository.PostCategory(*category)
}