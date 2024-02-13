package mocksandstructs

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

func MockInputCar() *domain.Car {
	return &domain.Car{
		ID:            "mockedID",
		Name:          "Mocked Car",
		Description:   "This is a mocked car for testing",
		DailyRate:     50.0,
		Available:     true,
		LicensePlate:  "ABC123",
		FineAmount:    10.0,
		Brand:         "Mocked Brand",
		CategoryID:    "mockedCategoryID",
		Specification: MockInputSpecifications(),
	}
}

func MockInputCars(numCars int) []*domain.Car {
	var cars []*domain.Car

	for i := 0; i < numCars; i++ {
		car := &domain.Car{
			ID:            fmt.Sprintf("mockedID%d", i),
			Name:          fmt.Sprintf("Mocked Car %d", i),
			Description:   fmt.Sprintf("This is a mocked car %d for testing", i),
			DailyRate:     50.0 + float64(i),
			Available:     true,
			LicensePlate:  fmt.Sprintf("ABC%d123", i),
			FineAmount:    10.0 + float64(i),
			Brand:         fmt.Sprintf("Mocked Brand %d", i),
			CategoryID:    fmt.Sprintf("mockedCategoryID%d", i),
			Specification: MockInputSpecifications(),
		}

		cars = append(cars, car)
	}

	return cars
}

func MockInputSpecification() *domain.Specification {
	return &domain.Specification{
		ID:          "mockedSpecID",
		Name:        "Mocked Specification",
		Description: "This is a mocked specification for testing",
		CarID:       "mockedCarID",
	}
}

func MockInputSpecifications() []*domain.Specification {
	return []*domain.Specification{
		{
			ID:          "mockedSpecID",
			Name:        "Mocked Specification",
			Description: "This is a mocked specification for testing",
			CarID:       "mockedCarID",
		},
		// Quando precisarmos, só colocar aqui mais especifications
	}
}
