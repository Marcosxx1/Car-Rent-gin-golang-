package usecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/category-controller/category-dtos"
	"github.com/rs/xid"
)

func PostCategoryUseCase(registerCategory categorydtos.CategoryInputDTO, categoryRepository repositories.CategoryRepository) (*categorydtos.CategoryOutputDTO, error) {
	existingCategory, err := categoryRepository.FindCategoryByName(registerCategory.Name)
	if err != nil {
			return nil, err
	}

	if existingCategory != nil {
			return nil, errors.New("category already exists")
	}

	newCategory := &domain.Category{
			ID:          xid.New().String(),
			Name:        registerCategory.Name,
			Description: registerCategory.Description,
	}

	if err := categoryRepository.PostCategory(newCategory); err != nil {
			return nil, fmt.Errorf("failed to create category: %w", err)
	}

	outputDTO := &categorydtos.CategoryOutputDTO{
			ID:          newCategory.ID,
			Name:        newCategory.Name,
			Description: newCategory.Description,
			CreatedAt:   newCategory.CreatedAt,
	}

	return outputDTO, nil
}