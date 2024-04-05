package ordercontroller

import (
	"fmt"
	"net/http"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetOrderController retrieves an order by order ID or user ID.
// @Summary Retrieve an order
// @Description Retrieve an order by order ID or user ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param options query string false "Order ID"
// @Success 200 {object} orderdto.OrderOutputDTO "Retrieved order"
// @Router /orders [get]

func GetOrderByQueryController(context *gin.Context, getOrderByQueryUseCase *orderusecases.GetOrderByQueryUseCase) {
	options := context.Query("options")
	userID := context.Query("userID")

	if options == "" && userID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Either options or userID must be provided"})
		return
	}
	//??????????????????????????????????????????????????????????????????????????
	var searchOptions *orderdto.OrderInputCompleteDTO
	switch options {
	case "id":
		searchOptions.CarID = options
	case "user_id":
		searchOptions.UserID = options
	case "car_id":
		searchOptions.CarID = options
	case "rental_start_date":
		startDate, err := utils.StringToTime(options)
		if err != nil {
			fmt.Println("Algo")
		}
		searchOptions.RentalStartDate = startDate
	case "rental_end_date":
		endDate, err := utils.StringToTime(options)
		if err != nil {
			fmt.Println("Algo")
		}
		searchOptions.RentalEndDate = endDate
	case "total_cost":
		parsedTotalCost, err := utils.StringToFloat64(options)
		if err != nil {
			fmt.Println("total cost error")
		}
		searchOptions.TotalCost = parsedTotalCost
	default:
		// throw error?
	}
	fmt.Println(searchOptions)
	order, err := getOrderByQueryUseCase.Execute(searchOptions)
	if err != nil {
		validation_errors.NewError(context, http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, order)
}
