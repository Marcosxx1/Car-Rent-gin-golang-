package mocksandstructs

import (
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller/specification-dtos"
)

func MockReturningCar() *dtos.CarOutputDTO {
	return &dtos.CarOutputDTO{
		ID:            "mockedID",
		Name:          "Mocked Car",
		Description:   "This is a mocked car for testing",
		DailyRate:     50.0,
		Available:     true,
		LicensePlate:  "ABC123",
		FineAmount:    10.0,
		Brand:         "Mocked Brand",
		CategoryID:    "mockedCategoryID",
		Specification: MockReturningSpecifications(),
	}
}

func MockReturningSpecifications() []*specificationdtos.SpecificationOutputDto {
	return []*specificationdtos.SpecificationOutputDto{
		{
			ID:          "spec1",
			Name:        "Mocked Specification 1",
			Description: "Description for Mocked Specification 1",
			CarID:       "mockedID",
		},
		{
			ID:          "spec2",
			Name:        "Mocked Specification 2",
			Description: "Description for Mocked Specification 2",
			CarID:       "mockedID",
		},
	}
}
