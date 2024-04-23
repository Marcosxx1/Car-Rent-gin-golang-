package validation_errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// NewValidationError trata erros de validação e responde ao cliente com mensagens formatadas de erro.
// Se o erro for do tipo validator.ValidationErrors, ele formatará as mensagens de erro de validação.
// Caso contrário, usará a função NewError para tratar o erro.
func NewValidationError(context *gin.Context, status int, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		context.JSON(status, formatValidationErrors(errs))
		return
	}
	NewError(context, status, err)
}

// formatValidationErrors formata mensagens de erro para erros de validação.
// Ele itera sobre os erros de validação e constrói mensagens de erro formatadas.
func formatValidationErrors(errs validator.ValidationErrors) gin.H {
	var errorMsgs []string
	for _, e := range errs {
		errorMsgs = append(errorMsgs, fmt.Sprintf("%s is required", e.Field()))
	}
	return gin.H{
		"errors": gin.H{
			"message": errorMsgs,
		},
	}
}

// NewErrors trata erros gerais e responde ao cliente com mensagens de erro padrão.
// Cria uma nova instância de HTTPErrors com o código de status e a mensagem de erro.
func NewErrors(context *gin.Context, status int, err error) {
	er := HTTPErrors{
		Code:    status,
		Message: err.Error(),
	}
	context.JSON(status, er)
}

// HTTPErrors é uma estrutura para representar erros HTTP com código e mensagem.
type HTTPErrors struct {
	Code    int    `json:"code" example:"400"`    // Código de status HTTP.
	Message string `json:"message" example:"status bad request"` // Mensagem de erro.
}
