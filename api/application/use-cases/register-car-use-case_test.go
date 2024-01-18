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

func (m *MockCarRepository) FindCarByLicensePlate(licensePlate string) (*domain.Car, error) {
	args := m.Called(licensePlate)
	return nil, args.Error(1) 
}


func (m *MockCarRepository) DeleteCar(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCarRepository) FindCarById(id string) (*domain.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Car), args.Error(1)
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
	mockRepo.On("FindCarByLicensePlate", mock.Anything).Return(nil, nil).Times(0)

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

	invalidRequest := usecases.RegisterCarRequest{}
 
	mockRepo.On("FindCarByLicensePlate", mock.Anything).Return(nil, nil).Times(0)
	resultingCar, err := usecases.RegisterCarUseCase(invalidRequest, mockRepo)

	assert.Error(t, err)
	assert.Nil(t, resultingCar)
	assert.Contains(t, err.Error(), "name is required")  
	mockRepo.AssertExpectations(t)
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
		Name:         "Missing Required Fields Car",
		Description:  "Missing Required Fields Description",
		DailyRate:    0.0,
		Available:    false,
		LicensePlate: "",
		FineAmount:   0.0,
		Brand:        "Missing Required Fields Brand",
		CategoryId:   "789",
		CreatedAt:    time.Now(),
	}

	invalidFieldValues := domain.Car{
		DailyRate:    -10.0,
		LicensePlate: "",
		FineAmount:   -5.0,
	}

	mixedData := domain.Car{
		Name:         "Mixed Data Car",
		Description:  "Mixed Data Description",
		DailyRate:    -10.0,
		Available:    true,
		LicensePlate: "",
		FineAmount:   10.0,
		Brand:        "Mixed Data Brand",
		CategoryId:   "987",
		CreatedAt:    time.Now(),
	}

	boundaryValues := domain.Car{
		// boundary values...
		DailyRate:  0.0,
		FineAmount: 0.0,
	}

	testCases := []struct {
		car       domain.Car
		isValid   bool
		errorText string
	}{
		{validData, true, ""},
		{missingRequiredFields, false, "required"},
		{invalidFieldValues, false, "required"},
		{mixedData, false, "required"},
		{boundaryValues, false, "required"},
	}

	for _, tc := range testCases {
		t.Run(tc.errorText, func(t *testing.T) {
			err := error_handling.ValidateStruct(tc.car)

			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.errorText)

				if err != nil {
					assert.Contains(t, err.Error(), tc.errorText)
				}
			}
		})
	}

}

/* func TestFindCarByLicensePlate(t *testing.T) {
	mockRepo := new(MockCarRepository)

	existingLicensePlate := "ABC123"
	nonExistingLicensePlate := "XYZ789"

	expectedExistingCar := &domain.Car{
		Id:           xid.New().String(),
		Name:         "Existing Car",
		Description:  "Existing Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: existingLicensePlate,
		FineAmount:   10.0,
		Brand:        "Existing Brand",
		CategoryId:   "456",
		CreatedAt:    time.Now(),
	}

	expectedNonExistingCar := (*domain.Car)(nil) 

	mockRepo.On("FindCarByLicensePlate", existingLicensePlate).Return(expectedExistingCar, nil)
	mockRepo.On("FindCarByLicensePlate", nonExistingLicensePlate).Return(expectedNonExistingCar, nil)

	existingCar, err := mockRepo.FindCarByLicensePlate(existingLicensePlate)
	assert.NoError(t, err)
	assert.NotNil(t, existingCar)
	assert.Equal(t, expectedExistingCar, existingCar)

	nonExistingCar, err := mockRepo.FindCarByLicensePlate(nonExistingLicensePlate)
	assert.NoError(t, err)
	assert.Nil(t, nonExistingCar)

	mockRepo.AssertExpectations(t)
}
 */