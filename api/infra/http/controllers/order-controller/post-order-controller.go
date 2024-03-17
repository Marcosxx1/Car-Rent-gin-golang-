package orderendpoints

import (
	"net/http"

	orderusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/order-use-cases"
	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/order-controller/order-dto"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

func PostOrderController(context *gin.Context, postOrderUseCase *orderusecases.PostOrderUseCase) {

	userID := context.GetString("user_id")     //comes from the auth middleware
	unsanitizedCarId := context.Param("carID") // comes from path

	carId := middlewares.SanitizeString(unsanitizedCarId) // sanitize

	var partialRequest orderdto.OrderInputPartialDTO
	if err := context.ShouldBindJSON(&partialRequest); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	var request orderdto.OrderInputCompleteDTO
	request.UserID = userID
	request.CarID = carId
	request.RentalStartDate = partialRequest.RentalStartDate
	request.RentalEndDate = partialRequest.RentalEndDate
	request.TotalCost = partialRequest.TotalCost
	request.OrderStatus = partialRequest.OrderStatus

	createdOrder, err := postOrderUseCase.Execute(request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdOrder)
}
