package ordercontroller

import (
	"net/http"

	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	"github.com/gin-gonic/gin"
)

// DeleteOrderController deletes an order by ID.
// @Summary Delete order by ID
// @Description Deletes an order by its ID
// @ID delete-order-by-id
// @Tags Orders
// @Accept json
// @Produce json
// @Security 		BearerAuth
// @Param order_id path string true "Order ID"
// @failure      400              {string}  string    "error"
// @response     default          {string}  string    "other error"
// @Failure 404 {string} asd "Order not found"
// @Failure 500 {string} string "Internal server error"
// @Router /orders/{order_id} [delete]
func DeleteOrderController(context *gin.Context, deleteOrderUseCase *orderusecases.DeleteOrderUseCase) {
	orderId := context.Param("order_id")

	err := deleteOrderUseCase.Execute(orderId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, nil)
}
