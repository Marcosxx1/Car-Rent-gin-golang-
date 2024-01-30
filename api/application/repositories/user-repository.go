package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type UserRepository interface {
	PostUser(userData *domain.User) error
	GetById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(id string, data *domain.User) error
}
