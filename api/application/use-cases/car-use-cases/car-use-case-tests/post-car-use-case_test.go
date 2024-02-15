package usecases_test

import (
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases/car-use-case-tests/mocks-and-structs"
	s "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/assert"
)

func TestPostCar_Success(t *testing.T) {

	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewPostCarUseCase(mockCarRepo, mockSpecRepo)

	inputDTOS := m.MockIncomingCarForCreation()

	// Set up expectations for the mock repositories
	mockCarRepo.On("FindCarByLicensePlate", "ABC123").Return(&domain.Car{}, nil) // Assume car does not exist
	mockCarRepo.On("RegisterCar").Return(nil)                                    // Assume car registration is successful
	mockSpecRepo.On("PostMultipleSpecifications").Return(nil)                    // Assume specification creation is successful

	// Execute the use case
	outputDTO, err := useCase.ExecuteConcurrently(inputDTOS)

	// Assert the results
	assert.NoError(t, err, "Unexpected error in PostCarUseCase.Execute")
	assert.NotNil(t, outputDTO, "OutputDTO should not be nil")

	// Add more assertions if needed based on your use case
	// ...

	// Verify that the expected methods were called on the mock repositories
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}
