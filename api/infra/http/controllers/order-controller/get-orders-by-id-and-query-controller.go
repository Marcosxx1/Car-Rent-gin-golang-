package ordercontroller

import (
	"net/http"
	"net/url"

	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetOrderByQueryController fetches orders based on query parameters.
//
// @Summary Get orders by query parameters
// @Description Fetches orders based on query parameters such as options and userID.
// @Tags Orders
// @Accept json
// @Produce json
// @Param options query string false "Query parameters for filtering orders"
// @Success 200 {array} orderdto.OrderOutputDTO "OK"
// @Router /orders [get]
func GetOrderByQueryController(context *gin.Context, getOrderByQueryUseCase *orderusecases.GetOrderByQueryUseCase) {
	options := context.Query("options")

	if options == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Either options or userID must be provided"})
		return
	}

	queryParams, err := url.ParseQuery(options)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse options query parameter"})
		return
	}

	order, err := getOrderByQueryUseCase.Execute(queryParams)
	if err != nil {
		validation_errors.NewError(context, http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, order)
}
