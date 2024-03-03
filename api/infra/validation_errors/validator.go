package validation_errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	fmt.Println(obj)

	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	var errorMessages []string

	switch e := err.(type) {
	case validator.ValidationErrors:
		for _, validationError := range e {
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
	case *validator.InvalidValidationError:
		// Handle the case where the error is of type *validator.InvalidValidationError
		// You might want to log or handle this differently based on your application needs.
		errorMessages = append(errorMessages, "Invalid validation error")
	default:
		return e // Return the original error if it's not one of the expected types.
	}

	return errors.New(strings.Join(errorMessages, "; "))
}
