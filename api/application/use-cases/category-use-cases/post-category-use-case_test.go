package categoryusecases

import (
	"errors"
	"testing"

	categorydtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/category"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPostCategoryFailedValidation(t *testing.T) {

	mockInputMissingBothFields := &categorydtos.CategoryInputDTO{
		Name:        "",
		Description: "",
	}
	runCategoryValidationTest(t, mockInputMissingBothFields, "name is required; description is required")

	mockInputMissingName := &categorydtos.CategoryInputDTO{
		Name:        "",
		Description: "description",
	}
	runCategoryValidationTest(t, mockInputMissingName, "name is required")

	mockInputMissingDescription := &categorydtos.CategoryInputDTO{
		Name:        "name",
		Description: "",
	}
	runCategoryValidationTest(t, mockInputMissingDescription, "description is required")
}

func runCategoryValidationTest(t *testing.T, request *categorydtos.CategoryInputDTO, expectedError string) {
	mockRepo := new(databasemocks.MockCategoryRepository)
	mockUseCase := NewPostCategoryUseCase(mockRepo)

	value, err := mockUseCase.Execute(request)

	if expectedError == "" {
		assert.Nil(t, err)
		assert.NotNil(t, value)
	} else {
		assert.NotNil(t, err)
		assert.Nil(t, value)
	}
}

func TestCreateCategoryAlreadyExists(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)
	mockUseCase := NewPostCategoryUseCase(mockRepo)

	mockRepo.On("FindCategoryByName", mock.AnythingOfType("string")).Return(&domain.Category{}, nil)

	mockInput := &categorydtos.CategoryInputDTO{
		Name:        "name",
		Description: "description",
	}

	value, err := mockUseCase.Execute(mockInput)

	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestErrorWhileFetchingData(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)
	mockUseCase := NewPostCategoryUseCase(mockRepo)

	mockRepo.On("FindCategoryByName", mock.AnythingOfType("string")).Return((&domain.Category{}), errors.New("some error"))

	mockInput := &categorydtos.CategoryInputDTO{
		Name:        "name",
		Description: "description",
	}

	value, err := mockUseCase.Execute(mockInput)

	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestFailedToCreateCategory(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)
	mockUseCase := NewPostCategoryUseCase(mockRepo)

	mockInput := &categorydtos.CategoryInputDTO{
			Name:        "name",
			Description: "description",
	}

	mockRepo.On("FindCategoryByName", mock.AnythingOfType("string")).Return((*domain.Category)(nil), nil)
	mockRepo.On("PostCategory", mock.AnythingOfType("*domain.Category")).Return(errors.New("failed to create category"))

	value, err := mockUseCase.Execute(mockInput)

	assert.Nil(t, value)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "failed to create category: failed to create category")
}

func TestSuccessfullyCategoryCreated(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)
	mockUseCase := NewPostCategoryUseCase(mockRepo)

	mockInput := &categorydtos.CategoryInputDTO{
			Name:        "name",
			Description: "description",
	}

	mockRepo.On("FindCategoryByName", mock.AnythingOfType("string")).Return((*domain.Category)(nil), nil)
	mockRepo.On("PostCategory", mock.AnythingOfType("*domain.Category")).Return(nil)

	value, err := mockUseCase.Execute(mockInput)

	assert.NotNil(t, value)
	assert.Nil(t, err)
}