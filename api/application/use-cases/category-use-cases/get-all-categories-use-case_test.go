package categoryusecases

import (
	"errors"
	"testing"

	testutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/test-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)

var mockCategories []*domain.Category
var mockEmptyCategories []*domain.Category

func init() {
	mockCategories = testutils.GenerateDomainCategories(2)
}

func TestGetAllCategoriesSuccess(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)

	mockRepo.On("GetAll").Return(mockCategories, nil)

	useCase := NewGetAllCategoriesUseCase(mockRepo)

	output, err := useCase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(output), len(mockCategories))

}

func TestGetAllCategoriesErrorRetrievingCategories(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)

	mockRepo.On("GetAll").Return(([]*domain.Category)(nil), errors.New("error retrieving categories"))

	usecase := NewGetAllCategoriesUseCase(mockRepo)

	output, err := usecase.Execute()

	assert.Nil(t, output, "Expected output to be nil, it wasn't")
	assert.NotNil(t, err, "Expected err to not be nil, it was") // almost

	assert.Error(t, err, "Expected error, but no error got returned")
}

func TestGetAllCategoriesNoCategoryFound(t *testing.T) {
	mockRepo := new(databasemocks.MockCategoryRepository)

	mockRepo.On("GetAll").Return(mockEmptyCategories, nil)

	usecase := NewGetAllCategoriesUseCase(mockRepo)

	output, err := usecase.Execute()

	assert.Nil(t, output, "Expected output to be nil, it wasn't")
	assert.NotNil(t, err, "Expected err to not be nil, it was") // almost
	assert.Contains(t, err.Error(), "no categories found", "Expected error message to contain 'no categories found'")
	assert.Error(t, err, "Expected error, but no error got returned")
}
