package usecases_test

import (
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) RegisterCar(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockCarRepository) FindCarByLicensePlate(licensePlate string) (*domain.Car, error) {
	args := m.Called(licensePlate)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Car), args.Error(1)
}


func (m *MockCarRepository) FindCarById(id string) (*domain.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func TestPostCarUseCase_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockCarRepository)

	// Set up expectations for FindCarByLicensePlate
	mockRepo.On("FindCarByLicensePlate", mock.AnythingOfType("string")).
		Return(nil, nil)

	mockRepo.On("RegisterCar", mock.AnythingOfType("*domain.Car")).
		Return(nil)


	registerRequest := dtos.CarInputDTO{
		Name:         "Car Model XYZ",
		Description:  "A description of the car",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "MAS12345",
		FineAmount:   10.0,
		Brand:        "CarBrand",
	}

	result, err := usecases.PostCarUseCase(registerRequest, mockRepo)

	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, result, "Expected a non-nil result")

	mockRepo.AssertExpectations(t)
}

func TestPostCarUseCase_ValidationFailure(t *testing.T) {
	mockRepo := new(MockCarRepository)

	invalidRequest := dtos.CarInputDTO{}

	mockRepo.On("FindCarByLicensePlate", mock.Anything).Return(nil, nil).Times(0)
	resultingCar, err := usecases.PostCarUseCase(invalidRequest, mockRepo)

	assert.Error(t, err)
	assert.Nil(t, resultingCar)
	assert.Contains(t, err.Error(), " is required")
	mockRepo.AssertExpectations(t)
}

func TestCarValidation(t *testing.T) {
	validData := domain.Car{
		ID:           "123",
		Name:         "Valid Car",
		Description:  "Valid Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   10.0,
		Brand:        "Valid Brand",
	}

	missingRequiredFields := domain.Car{
		Name:         "Missing Required Fields Car",
		Description:  "Missing Required Fields Description",
		DailyRate:    0.0,
		Available:    false,
		LicensePlate: "",
		FineAmount:   0.0,
		Brand:        "Missing Required Fields Brand",
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
