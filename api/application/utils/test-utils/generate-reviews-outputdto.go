package testutils

import (
	"fmt"
	"math/rand"

	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/review"
)

func GenerateReviews(numReviews int) []*reviewdto.ReviewOutputDTO {

	var reviews []*reviewdto.ReviewOutputDTO

	for i := 0; i < numReviews; i++ {
		id := fmt.Sprintf("%d", i+1)
		userID := fmt.Sprintf("user%d", rand.Intn(1000))
		carID := fmt.Sprintf("car%d", rand.Intn(1000))
		rating := rand.Intn(5) + 1
		content := fmt.Sprintf("Review for car %s by user %s", carID, userID)

		review := &reviewdto.ReviewOutputDTO{
			ID:      id,
			UserId:  userID,
			CarId:   carID,
			Rating:  &rating,
			Content: content,
		}

		reviews = append(reviews, review)
	}

	return reviews

}
