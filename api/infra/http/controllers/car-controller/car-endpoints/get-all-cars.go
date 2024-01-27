package carendpoints

import (
	"strconv"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/gin-gonic/gin"
)

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
