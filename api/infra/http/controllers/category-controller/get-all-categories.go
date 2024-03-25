package categorycontroller

import (
	"net/http"

	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/category"
	categoryusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/category-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// ListCategoriesController handles the HTTP GET request to retrieve a list of categories.
//
//	@Summary		Retrieve a list of categories
//	@Description	Retrieve a list of all categories.
//	@ID				list-categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int								false	"Limit the number of categories to be retrieved"
//	@Param			offset	query		int								false	"Offset for pagination of categories"
//	@Success		200		{array}		categorydtos.CategoryOutputDTO	"List of categories"
//	@Failure		400		{object}	validation_errors.HTTPError
//	@Router			/api/v1/category/list [get]
func ListCategoriesController(context *gin.Context, listOfCategories *categoryusecases.GetAllCategoriesUseCase) {
	var listOfCategoriesToBeReturned []*categorydtos.CategoryOutputDTO

	listOfCategoriesToBeReturned, err := listOfCategories.Execute()
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, listOfCategoriesToBeReturned)
}
