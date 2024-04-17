package databasemocks

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

type MockReviewRepository struct {
	mock.Mock
}

func (m *MockReviewRepository) CreateReview(review *domain.Reviews) (*domain.Reviews, error) {
	args := m.Called(review)

	return args.Get(0).(*domain.Reviews), args.Error(1)
}

func (m *MockReviewRepository) GetReviewByID(reviewId string) (*domain.Reviews, error) {
	args := m.Called(reviewId)

	return args.Get(0).(*domain.Reviews), args.Error(1)
}

func (m *MockReviewRepository) GetAllReviews() ([]*domain.Reviews, error) {
	args := m.Called()

	return args.Get(0).([]*domain.Reviews), args.Error(1)
}

func (m *MockReviewRepository) DeleteReview(reviewID string) error {
	args := m.Called(reviewID)

	return args.Error(0)
}

func (m *MockReviewRepository) UpdateReview(id string, review *domain.Reviews) error {
	args := m.Called(id, review)

	return args.Error(0)
}
