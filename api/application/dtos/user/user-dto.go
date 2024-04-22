package userdtos

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/go-playground/validator/v10"
)

type UserInputDTO struct {
	Name     string      `json:"name" validate:"required"`
	Email    string      `json:"email" validate:"required,email"`
	Password string      `json:"password" validate:"required,min=8"`
	Role     domain.Role `json:"role" validate:"required,oneof=admin user manager"`
	Status   bool        `json:"status" validate:"required"`
	Avatar   string      `json:"avatar"`
}

type UserUpdateDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name" validte:"required"`
	Email  string `json:"email" validate:"required,email"`
	Status bool   `json:"status" validte:"required"`
	Avatar string `json:"avatar"`
}

type UserOutPutDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    bool   `json:"status"`
	Avatar    string `json:"avatar"` 
	CreatedAt string `json:"created_at"`
}

func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	return role == string(domain.RoleAdmin) || role == string(domain.RoleUser)
}
