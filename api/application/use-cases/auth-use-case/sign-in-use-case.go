package authusecase

import (
	"fmt"
	"sync"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases/user-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/hash-password"
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

func (useCase *PostUserUseCase) Execute(userInputDto userdtos.UserInputDTO) (*userdtos.UserOutPutDTO, error) {
	resultChan := make(chan *userdtos.UserOutPutDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(3)
	go userutils.UserValidation(&wg, errorChan, validationErrorSignal, userInputDto)
	go useCase.performUserCreation(&wg, errorChan, validationErrorSignal, resultChan, userInputDto)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
		close(validationErrorSignal)
	}()

	for {
		select {
		case result := <-resultChan:
			return result, nil
		case err := <-errorChan:
			return nil, err
		}
	}
}

func (useCase *PostUserUseCase) performUserCreation(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, resultChan chan<- *userdtos.UserOutPutDTO, userInputDto userdtos.UserInputDTO) {
	defer wg.Done()

	hashedPassword, err := hashpassword.HashPassword(userInputDto.Password)
	if err != nil {
		errorChan <- fmt.Errorf("failed to hash password: %w", err)
		validationErrorSignal <- true
		return
	}

	newUser := &domain.User{
		ID:       xid.New().String(),
		Name:     userInputDto.Name,
		Email:    userInputDto.Email,
		Password: hashedPassword,
		Role:     userInputDto.Role,
		Status:   userInputDto.Status,
		Avatar:   userInputDto.Avatar,
	}

	if err := useCase.userRepository.PostUser(newUser); err != nil {
		errorChan <- fmt.Errorf("failed to save user: %w", err)
		validationErrorSignal <- true
		return
	}

	resultChan <- &userdtos.UserOutPutDTO{
		ID:     newUser.ID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Status: newUser.Status,
		Avatar: newUser.Avatar,
	}

	validationErrorSignal <- false
}
