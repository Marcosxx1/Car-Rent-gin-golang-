package error_handling

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	var errorMessages []string

	validationErrors := err.(validator.ValidationErrors)
	for _, validationError := range validationErrors {
		field := strings.ToLower(validationError.StructField())

		switch validationError.Tag() {
		case "required":
			errorMessages = append(errorMessages, field+" is required")
		case "max":
			errorMessages = append(errorMessages, field+" is required with max "+validationError.Param())
		case "min":
			errorMessages = append(errorMessages, field+" is required with min "+validationError.Param())
		case "email":
			errorMessages = append(errorMessages, field+" is invalid")
		}
	}

	return errors.New(strings.Join(errorMessages, "; "))
}
