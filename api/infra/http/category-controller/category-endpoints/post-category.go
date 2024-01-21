package categoryendpoints

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/category-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/category-controller/category-dtos"
	"github.com/gin-gonic/gin"
)

func PostCategoryController(context *gin.Context) {
	categoryRepository := database.PGCategory{}
 
	var request categorydtos.CategoryInputDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := categorydtos.CategoryInputDTO{
		Name:        request.Name,
		Description: request.Description,
	}

	createdCategory, err := usecases.PostCategoryUseCase(category, &categoryRepository)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusCreated, createdCategory)
	}
} 
