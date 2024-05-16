package specificationcontroller

import (
	"net/http"

	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/specification"
	specificationusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// PostSpecificationController handles the HTTP POST request to create a new specification.
// @Summary     Create a new specification
// @Description Create a new specification with the provided information.
// @ID          post-specification
// @Tags        Specification
// @Accept      json
// @Produce     json
// @Security 	BearerAuth
// @Param       request   	  body     specificationdtos.SpecificationInputDto  true "Specification information to be created"
// @Success	    201   		  {object} specificationdtos.SpecificationOutputDto "Successfully created specification"
// @Failure			400       {object} validation_errors.HTTPError
// @Router			/api/v1/specification/create [post]
func PostSpecificationController(context *gin.Context, specificationusecases *specificationusecases.PostSpecificationUseCase) {

	var request *specificationdtos.SpecificationInputDto

	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}

	createdSpecification, err := specificationusecases.Execute(request)
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusCreated, createdSpecification)
}
