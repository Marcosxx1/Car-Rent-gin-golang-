package categoryendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/category-use-cases"
	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/category-controller/category-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// PostCategoryController handles the HTTP POST request to create a new category.
// @Summary     Create a new category
// @Description Create a new category with the provided information.
// @ID          post-category
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       request    body    categorydtos.CategoryInputDTO true "Category information to be created"
// @Success	    201   		{object} categorydtos.CategoryOutputDTO "Successfully created category"
// @Failure			400       {object} validation_errors.HTTPError
// @Router			/api/v1/category/create [post]
func PostCategoryController(context *gin.Context, categoryRepository repositories.CategoryRepository) {

	var request categorydtos.CategoryInputDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}

	category := categorydtos.CategoryInputDTO{
		Name:        request.Name,
		Description: request.Description,
	}

	createdCategory, err := usecases.PostCategoryUseCase(category, categoryRepository)
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusCreated, createdCategory)
}
