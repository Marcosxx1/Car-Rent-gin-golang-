package usecases_test

import (
	"testing"
	"time"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain/error_handling"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) RegisterCar(car domain.Car) *domain.Car {
	args := m.Called(car)
	return args.Get(0).(*domain.Car)
}

func TestRegisterCarUseCase(t *testing.T) {
	mockRepo := new(MockCarRepository)

	registerRequest := usecases.RegisterCarRequest{
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   10.0,
		Brand:        "Test Brand",
		CategoryId:   "123",
	}
	
	expectedCar := &domain.Car{
		Id:           xid.New().String(),
 		Name:         registerRequest.Name,
		Description:  registerRequest.Description,
		DailyRate:    registerRequest.DailyRate,
		Available:    registerRequest.Available,
		LicensePlate: registerRequest.LicensePlate,
		FineAmount:   registerRequest.FineAmount,
		Brand:        registerRequest.Brand,
		CategoryId:   registerRequest.CategoryId,
 	}
	
	mockRepo.On("RegisterCar", mock.Anything).Return(expectedCar)

	resultingCar, err := usecases.RegisterCarUseCase(registerRequest, mockRepo)
  
	assert.NoError(t, err)
	assert.NotNil(t, resultingCar)
	assert.Equal(t, expectedCar, resultingCar) 
	assert.NotEmpty(t, resultingCar.Id)
	assert.NotEqual(t, "generated_id", resultingCar.Id)
 
	mockRepo.AssertExpectations(t)
}

func TestRegisterCarUseCase_ValidationFailure(t *testing.T) {
	mockRepo := new(MockCarRepository)

	invalidRequest := usecases.RegisterCarRequest{
	}

	resultingCar, err := usecases.RegisterCarUseCase(invalidRequest, mockRepo)

	assert.Error(t, err)
	assert.Nil(t, resultingCar)
	assert.Contains(t, err.Error(), "name is required") 
	assert.Empty(t, mockRepo.Calls)
}

func TestCarValidation(t *testing.T) {
	validData := domain.Car{
		Id:           "123",
		Name:         "Valid Car",
		Description:  "Valid Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   10.0,
		Brand:        "Valid Brand",
		CategoryId:   "456",
		CreatedAt:    time.Now(),
	}

	missingRequiredFields := domain.Car{
		// missing required fields...
	}

	invalidFieldValues := domain.Car{
		// invalid field values...
		DailyRate:    -10.0,
		LicensePlate: "",
		FineAmount:   -5.0,
	}

	mixedData := domain.Car{
		// valid and invalid mixed data...
		Name:         "Valid Car",
		Description:  "Valid Description",
		DailyRate:    -10.0,
		Available:    true,
		LicensePlate: "",
		FineAmount:   10.0,
		Brand:        "Valid Brand",
		CategoryId:   "456",
		CreatedAt:    time.Now(),
	}

	boundaryValues := domain.Car{
		// boundary values...
		DailyRate:    0.0,
		FineAmount:   0.0,
	}

	testCases := []struct {
		car       domain.Car
		isValid   bool
		errorText string
	}{
		{validData, true, ""},
		{missingRequiredFields, false, "required"},
		{invalidFieldValues, false, "name is required"},
		{mixedData, false, "dailyrate validation failed with tag: gte"},
		{boundaryValues, false, "name is required"},
	}

	for _, tc := range testCases {
		t.Run(tc.errorText, func(t *testing.T) {
			err := error_handling.ValidateStruct(tc.car)

			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errorText)
			}
		})
	}
}
