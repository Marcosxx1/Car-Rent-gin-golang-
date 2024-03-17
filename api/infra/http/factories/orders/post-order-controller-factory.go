package factory

import (
	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	orderendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/order-controller/order-endpoints"
	"github.com/gin-gonic/gin"
)

func PostOrderFactoryController(context *gin.Context) {
	userRepository := database.NewPGUserRepository()
	maintenanceRepository := database.NewPgMaintenanceRepository()
	orderRepository := database.NewPGOrderRepository()

	postOrderUseCase := orderusecases.NewPostOrderUseCase(userRepository, maintenanceRepository, orderRepository)

	orderendpoints.PostOrderController(context, postOrderUseCase)

}
