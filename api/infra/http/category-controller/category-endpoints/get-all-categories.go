package categoryendpoints

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/category-use-cases"
	"github.com/gin-gonic/gin"
)

 
func ListCategoriesController(context *gin.Context, categoryRepository repositories.CategoryRepository) {
/* 	limitStr := context.DefaultQuery("limit", "1")
	offsetStr := context.DefaultQuery("offset", "10") */

/* 	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	offset, errs := strconv.Atoi(offsetStr)
	if errs != nil {
		context.JSON(400, gin.H{
			"error": errs.Error(),
		})
		return
	} */

	listOfCategories, err := usecases.GetAllCategoriesUseCase(context, categoryRepository/* , limit, offset */)
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, listOfCategories)
}
