package testutils

import (
	"fmt"
	"math/rand"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

func GenerateDomainReviews(numReviews int) []*domain.Reviews {

	var reviews []*domain.Reviews

	for i := 0; i < numReviews; i++ {
		id := fmt.Sprintf("id%d", i)
		userId := fmt.Sprintf("user%d", i)
		carId := fmt.Sprintf("car%d", i)
		rating := rand.Intn(5) + 1

		review := &domain.Reviews{
			ID:     id,
			UserId: userId,
			CarId:  carId,
			Rating: &rating,
		}

		reviews = append(reviews, review)
	}

	return reviews
}
