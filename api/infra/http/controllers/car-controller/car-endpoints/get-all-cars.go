package carendpoints

import (
	"strconv"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/gin-gonic/gin"
)

// @Summary List all cars
// @Description Retrieve a list of cars with pagination support.
// @ID list-car
// @Tags Car
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Number of items per page (default is 10)"
// @Success 200 {array} dtos.CarOutputDTO "List of cars"
// @Failure				422					{array}		validation_errors.HTTPError
// @Router /api/v1/cars [get]
func ListCarController(context *gin.Context, carRepository repositories.CarRepository) {
	page, pageSize := getPaginationParameters(context)

	var wg sync.WaitGroup
	carChannel := make(chan []*dtos.CarOutputDTO)
	errChannel := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		listOfCars, err := usecases.GetAllCarsUseCase(carRepository, page, pageSize)
		if err != nil {
			errChannel <- err
			return
		}
		carChannel <- listOfCars
	}()

	go func() {
		wg.Wait()
		close(carChannel)
		close(errChannel)
	}()

	select {
	case listOfCars := <-carChannel:
		context.JSON(200, listOfCars)
	case err := <-errChannel:
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
}

func getPaginationParameters(context *gin.Context) (int, int) {
	page, err := strconv.Atoi(context.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(context.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	return page, pageSize
}
