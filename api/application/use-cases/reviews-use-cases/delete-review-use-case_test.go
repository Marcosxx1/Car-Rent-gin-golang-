package reviewusecases

import (
	"errors"
	"testing"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func TestDeleteReviewSuccess(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)

	// Mock GetReviewByID to return a non-nil review
	mockRepo.On("GetReviewByID", mock.AnythingOfType("string")).Return(&domain.Reviews{
		ID: "any_id",
	}, nil)

	// Mock DeleteReview to return nil (indicating success)
	mockRepo.On("DeleteReview", mock.AnythingOfType("string")).Return(nil)

	useCase := NewDeleteReviewUseCase(mockRepo)

	err := useCase.Execute("any_id")

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteReviewUseCase_GetReviewByIDError(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewDeleteReviewUseCase(mockRepo)

	mockRepo.On("GetReviewByID", mock.AnythingOfType("string")).Return((&domain.Reviews{}), errors.New("failed to get record: record not found"))

	err := useCase.Execute("any_id")

	assert.True(t, errors.Is(err, ErrGetRecord), "Expected ErrGetRecord error")
	mockRepo.AssertExpectations(t)
}

func TestDeleteReviewUseCase_RecordNotFound(t *testing.T){
	// Create a mock reviewRepository
	mockRepo := new(databasemocks.MockReviewRepository)

	// Initialize the DeleteReviewUseCase with the mock repository
	useCase := NewDeleteReviewUseCase(mockRepo)

	// Mock GetReviewByID to return a review with an empty ID
	mockRepo.On("GetReviewByID", "emptyID").Return(&domain.Reviews{ID: ""}, nil)

	// Execute the use case
	err := useCase.Execute("emptyID")

	// Check the error
	if err == nil || err.Error() != "record not found" {
		t.Errorf("Expected error: record not found, got: %v", err)
	}
}
func TestDeleteReviewUseCase_DeleteReviewError(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewDeleteReviewUseCase(mockRepo)

	mockRepo.On("GetReviewByID", mock.AnythingOfType("string")).Return(&domain.Reviews{
		ID: "any_id",
	}, nil)
	mockRepo.On("DeleteReview", mock.AnythingOfType("string")).Return(errors.New("failed to delete record"))

	err := useCase.Execute("any_id")

	assert.Contains(t, err.Error(), "failed to delete record")
	mockRepo.AssertExpectations(t)
}

/*
// este teste:
func TestDeleteReviewUseCase_GetReviewByIDError(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewDeleteReviewUseCase(mockRepo)

	mockRepo.On("GetReviewByID", mock.AnythingOfType("string")).Return((&domain.Reviews{}), errors.New("failed to get record: record not found"))

	err := useCase.Execute("any_id")

	assert.Equal(t, ErrGetRecord, err)
	mockRepo.AssertExpectations(t)
}

dá esse erro:

Error:      	Not equal:
            	expected: *errors.errorString(&errors.errorString{s:"failed to get record: record not found"})
            	actual  : *fmt.wrapError(&fmt.wrapError{msg:"failed to get record: record not found: failed to get record: record not found", err:(*errors.errorString)(0x77c290)})
Test:       	TestDeleteReviewUseCase_GetReviewByIDError

// por que:
 Quando utilizamos return fmt.Errorf em um teste, ele retorna um erro do tipo *fmt.wrapError.
 Isso significa que o erro está encapsulado dentro de um *fmt.wrapError.
 Isso pode causar problemas ao comparar os erros.

 Solução:

 Para comparar os erros, você pode utilizar os métodos errors.Is e errors.As.

*/

/*
// este teste:
func TestDeleteReviewSuccess(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)

	mockRepo.On("GetReviewByID", mock.AnythingOfType("string")).Return(&domain.Reviews{})
	mockRepo.On("DeleteReview",  mock.AnythingOfType("string")).Return(nil)

	useCase := NewDeleteReviewUseCase(mockRepo)

	err := useCase.Execute("any_id")

	assert.Nil(t, err)

}

// vai causar este erro:
// --- FAIL: TestDeleteReviewSuccess (0.00s)
// panic: assert: arguments: Cannot call Get(1) because there are 1 argument(s). [recovered]
// 	panic: assert: arguments: Cannot call Get(1) because there are 1 argument(s).

// Por causa dessa linha:
	mockRepo.On("GetReviewByID", mock.AnythingOfType("string")).Return(&domain.Reviews{})

	GetReviewByID retorna dois valores
	TODO(SEMPRE) - Sempre se atentar aos tipos de retornos ao testarmos

*/
