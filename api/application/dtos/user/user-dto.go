package userdtos

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type UserInputDTO struct {
	Name     string      `json:"name" binding:"required" validate:"required"`
	Email    string      `json:"email" binding:"required" validate:"required,email"`
	Password string      `json:"password" binding:"required" validate:"required,min=8"`
	Role     domain.Role `json:"role" binding:"required" validate:"required,oneof=admin user manager"`
	Status   bool        `json:"status" binding:"required" validate:"required"`
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

