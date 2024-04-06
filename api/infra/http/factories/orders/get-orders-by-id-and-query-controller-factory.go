package factory

import (
	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	ordercontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/order-controller"
	"github.com/gin-gonic/gin"
)

func GetOrderByQueryFactoryController(context *gin.Context) {
	userRepository := database.NewPGUserRepository()
	maintenanceRepository := database.NewPgMaintenanceRepository()
	orderRepository := database.NewPGOrderRepository()

	getOrderByQueryUseCase := orderusecases.NewGetOrderByQueryUseCase(userRepository, maintenanceRepository, orderRepository)

	ordercontroller.GetOrderByQueryController(context, getOrderByQueryUseCase)

}
