package userutils

import (
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

func ConvertUserDomainToOutputDTO(user *domain.User) *userdtos.UserOutPutDTO {
	return &userdtos.UserOutPutDTO{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
		Avatar: user.Avatar,
	}
}
