package mocksandstructs

/*
import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller/specification-dtos"
)

func MockCarFromDataBase() *domain.Car {
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
		Specification: MockSpecificationsFromDatabase(),
	}
}

func MockIncomingCarForCreation() *dtos.CarInputDTO {
	return &dtos.CarInputDTO{
		Name:          "Mocked Car",
		Description:   "This is a mocked car for testing",
		DailyRate:     50.0,
		Available:     true,
		LicensePlate:  "ABC123",
		FineAmount:    10.0,
		Brand:         "Mocked Brand",
		CategoryID:    "mockedCategoryID",
		Specification: MockIncomingSpecificationsToCreate(),
	}
}

func MockListOfCarsFromDatabase(numCars int) []*domain.Car {
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
			Specification: MockSpecificationsFromDatabase(),
		}

		cars = append(cars, car)
	}

	return cars
}

func MockSpecificationForCreation() *domain.Specification {
	return &domain.Specification{
		ID:          "mockedSpecID",
		Name:        "Mocked Specification",
		Description: "This is a mocked specification for testing",
		CarID:       "mockedCarID",
	}
}

func MockIncomingSpecificationForCreation() *specificationdtos.SpecificationInputDto {
	return &specificationdtos.SpecificationInputDto{
		Name:        "Mocked Specification",
		Description: "This is a mocked specification for testing",
		CarID:       "mockedCarID",
	}
}

func MockIncomingSpecificationsToCreate() []*specificationdtos.SpecificationInputDto {
	return []*specificationdtos.SpecificationInputDto{
		{
			Name:        "Mocked Specification",
			Description: "This is a mocked specification for testing",
			CarID:       "mockedCarID",
		},
		// Quando precisarmos, só colocar aqui mais especifications
	}
}

func MockSpecificationsFromDatabase() []*domain.Specification {
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
*/