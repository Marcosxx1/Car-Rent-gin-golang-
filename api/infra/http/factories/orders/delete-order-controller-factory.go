package factory

import (
	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	ordercontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/order-controller"
	"github.com/gin-gonic/gin"
)

func DeleteOrderFactoryController(context *gin.Context) {
	orderRepository := database.NewPGOrderRepository()

	deleteOrderUseCase := orderusecases.NewDeleteOrderUseCase(orderRepository)

	ordercontroller.DeleteOrderController(context, deleteOrderUseCase)

}
 