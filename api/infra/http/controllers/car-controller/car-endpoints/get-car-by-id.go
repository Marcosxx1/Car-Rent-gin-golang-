package carendpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/gin-gonic/gin"
)

// FindCarByIdController handles the HTTP GET request find a car.
// @Summary				Find car
// @Description		Find a car with the provided id
// @ID						get-car
// @Tags					Car
// @Accept				json
// @Produce				json
// @Param        	id   				path   		 string  true  "Car ID"
// @Success	    	201   			{object} 	dtos.CarOutputDTO " car"
// @Failure				422					{array}		validation_errors.HTTPError
// @Router				/api/v1/cars/{id} [get]
func FindCarByIdController(context *gin.Context,carRepository repositories.CarRepository) {

	id := context.Param("id")
	fmt.Printf("%+v\n", id)

	car, err := usecases.GetCarByIdUseCase(id, carRepository )
	if err != nil {
		log.Println("Error finding car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, car)
}