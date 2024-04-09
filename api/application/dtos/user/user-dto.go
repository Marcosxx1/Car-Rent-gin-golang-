package userdtos

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/go-playground/validator/v10"
)

type UserInputDTO struct {
	Name     string      `json:"name" binding:"required" validte:"required"`
	Email    string      `json:"email" binding:"required" validate:"required,email"`
	Password string      `json:"password" binding:"required" validate:"required,min=8"`
	Role     domain.Role `json:"role" binding:"required,oneof=admin user manager" validte:"required"`
	Status   bool        `json:"status" binding:"required" validte:"required"`
	Avatar   string      `json:"avatar"`
}

type UserUpdateDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name" binding:"required" validte:"required"`
	Email  string `json:"email" binding:"required" validate:"required,email"`
	Status bool   `json:"status" binding:"required" validte:"required"`
	Avatar string `json:"avatar" binding:"required"`
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
