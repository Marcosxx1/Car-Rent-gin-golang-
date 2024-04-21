package categoryusecases

import (
	"errors"

	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/category"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type GetAllCategoriesUseCase struct {
	categoryRepository repositories.CategoryRepository
}

func NewGetAllCategoriesUseCase(
	categoryRepository repositories.CategoryRepository) *GetAllCategoriesUseCase {
	return &GetAllCategoriesUseCase{
		categoryRepository: categoryRepository,
	}
}

func (useCase *GetAllCategoriesUseCase) Execute() ([]*categorydtos.CategoryOutputDTO, error) {

	allCategories, err := useCase.categoryRepository.GetAll()
	if err != nil {
		return nil, errors.New("error retrieving categories")
	}

	if len(allCategories) == 0 {
		return nil, errors.New("no categories found")
	}

	outPutDTO := make([]*categorydtos.CategoryOutputDTO, 0)
	for _, car := range allCategories {
		dto := &categorydtos.CategoryOutputDTO{
			ID:          car.ID,
			Name:        car.Name,
			Description: car.Description,
			CreatedAt:   car.CreatedAt,
		}
		outPutDTO = append(outPutDTO, dto)
	}

	return outPutDTO, nil
}
