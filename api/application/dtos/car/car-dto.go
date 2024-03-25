package cardtos

import (
	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/specification"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

/*

[]*domain.Specification : Declara um slice de ponteiros para objetos de domain.Specification
Cada elemento do slice é um objeto de domain.Specification

*[]domain.Specification. : Declara um ponteiro para um slice objetos de domain.Specification
O slice inteiro é envolto como ponteiro, e conseguimos a referencia para o próprio slice
Quando acessamos um elemento deste slice, acessamos uma instancia direta de domain.Specification

*[]*domain.Specification : Declara um ponteiro para um slice de ponteiros para objetos de domain.Specification

*/

type CarInputDTO struct {
	Name          string                                     `json:"name" validate:"required"`
	Description   string                                     `json:"description" validate:"required"`
	DailyRate     float64                                    `json:"daily_rate" validate:"required"`
	Available     bool                                       `json:"available" validate:"required"`
	LicensePlate  string                                     `json:"license_plate" validate:"required"`
	FineAmount    float64                                    `json:"fine_amount" validate:"required"`
	Brand         string                                     `json:"brand" validate:"required"`
	CategoryID    string                                     `json:"category_id" validate:"required"`
	Specification []*specificationdtos.SpecificationInputDto `json:"specification"`
}

type CarOutputDTO struct {
	ID            string                                      `json:"id"`
	Name          string                                      `json:"name"`
	Description   string                                      `json:"description"`
	DailyRate     float64                                     `json:"daily_rate"`
	Available     bool                                        `json:"available"`
	LicensePlate  string                                      `json:"license_plate"`
	FineAmount    float64                                     `json:"fine_amount"`
	Brand         string                                      `json:"brand"`
	CategoryID    string                                      `json:"category_id"`
	Specification []*specificationdtos.SpecificationOutputDto `json:"specification"`
}

func ConvertToOutputDTO(carID string, inputDTO *CarInputDTO) *CarOutputDTO {
	specificaDomain := utils.ConvertSpecificationToDomainCreate(inputDTO.Specification, carID)
	specificationOutPut := utils.ConvertSpecificationToDTO(specificaDomain)
	return &CarOutputDTO{
		ID:            carID,
		Name:          inputDTO.Name,
		Description:   inputDTO.Description,
		DailyRate:     inputDTO.DailyRate,
		Available:     inputDTO.Available,
		LicensePlate:  inputDTO.LicensePlate,
		FineAmount:    inputDTO.FineAmount,
		Brand:         inputDTO.Brand,
		CategoryID:    inputDTO.CategoryID,
		Specification: specificationOutPut,
	}
}

func ConvertDomainToOutPut(carID string, domainCarToBeConvertedToOutPut *domain.Car, specificationUpdated []*domain.Specification) *CarOutputDTO {

	return &CarOutputDTO{
		ID:            carID,
		Name:          domainCarToBeConvertedToOutPut.Name,
		Description:   domainCarToBeConvertedToOutPut.Description,
		DailyRate:     domainCarToBeConvertedToOutPut.DailyRate,
		Available:     domainCarToBeConvertedToOutPut.Available,
		LicensePlate:  domainCarToBeConvertedToOutPut.LicensePlate,
		FineAmount:    domainCarToBeConvertedToOutPut.FineAmount,
		Brand:         domainCarToBeConvertedToOutPut.Brand,
		CategoryID:    domainCarToBeConvertedToOutPut.CategoryID,
		Specification: utils.ConvertSpecificationToDTO(specificationUpdated),
	}
}
