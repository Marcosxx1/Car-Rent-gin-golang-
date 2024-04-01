package ordercontroller

import (
	"net/http"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
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

	// Convert options string to Options type
	var opt orderdto.Options
	switch options {
	case "id":
		opt = orderdto.ID
	case "user_id":
		opt = orderdto.UserID
	case "car_id":
		opt = orderdto.CarID
	case "rental_start_date":
		opt = orderdto.RentalStartDate
	case "rental_end_date":
		opt = orderdto.RentalEndDate
	case "total_cost":
		opt = orderdto.TotalCost
	default:
		opt = orderdto.Options(options)
	}

	order, err := getOrderByQueryUseCase.Execute(&opt)
	if err != nil {
		validation_errors.NewError(context, http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, order)
}
