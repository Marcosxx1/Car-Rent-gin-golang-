package authusecase

import (
	"fmt"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/hash-password"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/rs/xid"
)

type PostUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewPostUserUseCase(userRepository repositories.UserRepository) *PostUserUseCase {
	return &PostUserUseCase{
		userRepository: userRepository,
	}
}

func (useCase *PostUserUseCase) Execute(registerRequest userdtos.UserInputDTO) (*userdtos.UserOutPutDTO, error) {

	existUser, err := useCase.userRepository.FindByEmail(registerRequest.Email)

	if existUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", registerRequest.Email)
	}

	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashpassword.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &domain.User{
		ID:       xid.New().String(),
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: hashedPassword,
		Role:     registerRequest.Role,
		Status:   registerRequest.Status,
		Avatar:   registerRequest.Avatar,
	}

	if err := validation_errors.ValidateStruct(newUser); err != nil {
		return nil, err
	}

	if err := useCase.userRepository.PostUser(newUser); err != nil {
		return nil, fmt.Errorf("failed to singin user: %w", err)
	}

	outPut := &userdtos.UserOutPutDTO{
		ID:     newUser.ID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Status: newUser.Status,
		Avatar: newUser.Avatar,
	}

	return outPut, nil
}
