package validation_errors

import "github.com/gin-gonic/gin"

func NewError(context *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	context.JSON(status, er)
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError represents a custom error type for HTTP errors
type HTTPErrorCar struct {
	Code    int      `json:"code" example:"422"`
	Message string   `json:"message,omitempty" example:""`
	Errors  []string `json:"errors,omitempty" example:"['name is required', 'description is required', 'licenseplate is required', 'brand is required', 'A car needs a category to be registered']"`
}
type HTTPErrorMaintenance struct {
	Code    int      `json:"code" example:"422"`
	Message string   `json:"message,omitempty" example:""`
	Errors  []string `json:"errors,omitempty" example:"['name is required', 'description is required', 'licenseplate is required', 'brand is required', 'A car needs a category to be registered']"`
}


