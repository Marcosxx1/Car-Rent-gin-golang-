package usecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/category-controller/category-dtos"
	"github.com/gin-gonic/gin"
)

func GetAllCategoriesUseCase(context *gin.Context,
	categoryRepository repositories.CategoryRepository,
	/* 	limit int, offset int */) ([]*categorydtos.CategoryOutputDTO, error) {

	allCategories, err := categoryRepository.GetAll( /* limit, offset */ )
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
