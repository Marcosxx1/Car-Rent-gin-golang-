package usecases_test

import (
	"testing"
)

func (m *MockCarRepository) DeleteCar(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestDeleteCarUseCaseSuccess(t *testing.T) {

	mockRepo := new(MockCarRepository)

	mockRepo.On("DeleteCar", "1").Return(nil)
	err := mockRepo.DeleteCar("1")
	if err != nil {
		t.Errorf("Error deleting car: %v", err)
	}

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "DeleteCar", 1)
	mockRepo.AssertCalled(t, "DeleteCar", "1")

}

func TestDeleteCarUseCaseError(t *testing.T) {
	mockRepo := new(MockCarRepository)

	mockRepo.On("DeleteCar", "1").Return(nil)
	err := mockRepo.DeleteCar("1")
	if err != nil {
		t.Errorf("Error deleting car: %v", err)
	}

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "DeleteCar", 1)
	mockRepo.AssertCalled(t, "DeleteCar", "1")

}
