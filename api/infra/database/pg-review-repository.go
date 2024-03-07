package database

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGReviewRepository struct{}

func (repo *PGReviewRepository) CreateReview(review *domain.Reviews) (*domain.Reviews, error) {
	result := dbconfig.Postgres.Create(&review)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create review: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("failed to create review: no rows affected")
	}

	return review, nil
}

func (repo *PGReviewRepository) GetReviewByID(reviewID string) (*domain.Reviews, error) {
	var review domain.Reviews
	result := dbconfig.Postgres.Where("id = ?", reviewID).First(&review)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve review: %w", result.Error)
	}

	return &review, nil
}
