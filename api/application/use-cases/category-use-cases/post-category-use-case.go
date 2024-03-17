package categoryusecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/category-controller/category-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/rs/xid"
)

type PostCategoryUseCase struct {
	categoryRepository repositories.CategoryRepository
}

func NewPostCategoryUseCase(
	categoryRepository repositories.CategoryRepository) *PostCategoryUseCase {
	return &PostCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (useCase *PostCategoryUseCase) Execute(registerCategory *categorydtos.CategoryInputDTO) (*categorydtos.CategoryOutputDTO, error) {

	newCategory := &domain.Category{
		ID:          xid.New().String(),
		Name:        registerCategory.Name,
		Description: registerCategory.Description,
	}

	if err := validation_errors.ValidateStruct(newCategory); err != nil {
		return nil, err
	}

	existingCategory, err := useCase.categoryRepository.FindCategoryByName(registerCategory.Name)
	if err != nil {
		return nil, err
	}

	if existingCategory != nil {
		return nil, errors.New("category already exists")
	}

	if err := useCase.categoryRepository.PostCategory(newCategory); err != nil {
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
