package userutils

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
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
