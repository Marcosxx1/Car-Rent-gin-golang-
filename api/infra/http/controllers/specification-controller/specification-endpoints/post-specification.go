package specificationendpoints

import (
	"net/http"

	repo "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	specificationusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller/specification-dtos"
	"github.com/gin-gonic/gin"
)

func PostSpecificationController(context *gin.Context, specificationRepository repo.SpecificationRepository) {

	var request *specificationdtos.SpecificationInputDto

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	specification := &specificationdtos.SpecificationInputDto{
		Name:        request.Name,
		Description: request.Description,
		CarID:       request.CarID,
	}

	createdSpecification, err := specificationusecases.PostSpecificationUseCase(specification, specificationRepository)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, createdSpecification)
}
