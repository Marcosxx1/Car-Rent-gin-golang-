package specificationusecases

import (
	"errors"
	"testing"

	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/specification"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostSpecificationUseCase_Success(t *testing.T) {
	mockRepo := new(databasemocks.MockSpecificationRepository)

	useCase := NewPostSpecificationUseCase(mockRepo)

	mockRepo.On("FindSpecificationByName", mock.AnythingOfType("string")).Return((*domain.Specification)(nil), nil)

	mockRepo.On("PostSpecification", mock.AnythingOfType("*domain.Specification")).Return(nil)

	inputDto := &specificationdtos.SpecificationInputDto{
		Name:        "TestSpecification",
		Description: "Test Description",
		CarID:       "12345",
	}
	outputDto, err := useCase.Execute(inputDto)

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, outputDto, "Expected a non-nil outputDto")
	assert.Equal(t, inputDto.Name, outputDto.Name, "Expected names to match")

	mockRepo.AssertExpectations(t)
}

func TestPostSpecificationUseCase_SpecificationExists(t *testing.T) {
	mockRepo := new(databasemocks.MockSpecificationRepository)

	useCase := NewPostSpecificationUseCase(mockRepo)

	mockRepo.On("FindSpecificationByName", mock.AnythingOfType("string")).Return(&domain.Specification{}, nil)

	inputDto := &specificationdtos.SpecificationInputDto{
		Name:        "ExistingSpecification",
		Description: "Test Description",
		CarID:       "12345",
	}
	_, err := useCase.Execute(inputDto)

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, "specification already exists", err.Error(), "Expected specific error message")

	mockRepo.AssertExpectations(t)
}

func TestPostSpecificationUseCase_ErrorQueryingSpecification(t *testing.T) {
	mockRepo := new(databasemocks.MockSpecificationRepository)

	useCase := NewPostSpecificationUseCase(mockRepo)

	//(*domain.Specification)(nil) is a nil pointer to *domain.Specification !! REMEMBER
	mockRepo.On("FindSpecificationByName", mock.AnythingOfType("string")).Return((*domain.Specification)(nil), errors.New("error querying specification"))

	inputDto := &specificationdtos.SpecificationInputDto{
		Name:        "TestSpecification",
		Description: "Test Description",
		CarID:       "12345",
	}
	_, err := useCase.Execute(inputDto)

	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "error querying specification", "Expected error message to contain 'error querying specification'")

	mockRepo.AssertExpectations(t)
}

func TestPostSpecificationUseCase_ErrorCreatingSpecification(t *testing.T) {
	mockRepo := new(databasemocks.MockSpecificationRepository)

	useCase := NewPostSpecificationUseCase(mockRepo)

	mockRepo.On("FindSpecificationByName", mock.AnythingOfType("string")).Return((*domain.Specification)(nil), nil)

	mockRepo.On("PostSpecification", mock.AnythingOfType("*domain.Specification")).Return(errors.New("error creating specification"))

	inputDto := &specificationdtos.SpecificationInputDto{
		Name:        "TestSpecification",
		Description: "Test Description",
		CarID:       "12345",
	}
	_, err := useCase.Execute(inputDto)

	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "error creating specification", "Expected error message to contain 'error creating specification'")

	mockRepo.AssertExpectations(t)
}
