package ordercontroller

import (
	"net/http"

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
// @Param orderID query string false "Order ID"
// @Param userID query string false "User ID"
// @Success 200 {object} orderdto.OrderOutputDTO "Retrieved order"
// @Router /orders [get]
func GetOrderByIdAndQueryController(context *gin.Context, getOrderByIdAndQueryUseCase *orderusecases.GetOrderByIdAndQueryUseCase) {
	orderID := context.Query("orderID")

	if orderID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Either orderID or userID must be provided"})
		return
	} 

	order, err := getOrderByIdAndQueryUseCase.Execute(orderID)
	if err != nil {
		validation_errors.NewError(context, http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, order)
}
