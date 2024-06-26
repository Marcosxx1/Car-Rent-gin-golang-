package categorycontroller

import (
	"net/http"

	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/category"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/category-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// PostCategoryController handles the HTTP POST request to create a new category.
// @Summary     Create a new category (Authentication needed)
// @Description Create a new category with the provided information.
// @ID          post-category
// @Tags        Category
// @Accept      json
// @Produce     json
// @Security 	BearerAuth
// @Param       request    body    categorydtos.CategoryInputDTO true "Category information to be created"
// @Success	    201   		{object} categorydtos.CategoryOutputDTO "Successfully created category"
// @Router			/api/v1/category/create [post]
func PostCategoryController(context *gin.Context, categoryUseCase *usecases.PostCategoryUseCase) {

	var request categorydtos.CategoryInputDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}

	createdCategory, err := categoryUseCase.Execute(&request)
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusCreated, createdCategory)
}
