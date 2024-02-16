package dtos

import (
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller/specification-dtos"
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
	Description   string                                      `json:"description" validate:"required"`
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
	specificaDomain := repoutils.ConvertSpecificationToDomainCreate(inputDTO.Specification, carID)
	specificationOutPut := repoutils.ConvertSpecificationToDTO(specificaDomain)
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
