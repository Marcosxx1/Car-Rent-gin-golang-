package mocksandstructs

import (
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
