package categoryfactory

import (
	categoryusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/category-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	categorycontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/category-controller"
	"github.com/gin-gonic/gin"
)

func ListCategoriesControllerFactory(context *gin.Context) {
	categoryRepository := database.NewPGCategoryRepository()

	listOfCategories := categoryusecases.NewGetAllCategoriesUseCase(categoryRepository)

	categorycontroller.ListCategoriesController(context, listOfCategories)
}
