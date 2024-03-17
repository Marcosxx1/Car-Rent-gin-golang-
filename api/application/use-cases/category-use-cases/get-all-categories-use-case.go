package categoryusecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/category-controller/category-dtos"
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
		return nil, err
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
