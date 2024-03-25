package utils

import (
	specificationdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/specification"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/rs/xid"
)

func ConvertSpecificationToDTO(specificaion []*domain.Specification) []*specificationdtos.SpecificationOutputDto {
	var SpecificationsDTO []*specificationdtos.SpecificationOutputDto
	for _, spec := range specificaion {
		SpecificationsDTO = append(SpecificationsDTO, &specificationdtos.SpecificationOutputDto{
			ID:          spec.ID,
			Name:        spec.Name,
			Description: spec.Description,
			CarID:       spec.CarID,
		})
	}
	return SpecificationsDTO
}

func ConvertOutPutSpecificationToDomainUpdate(specification []*specificationdtos.SpecificationOutputDto) []*domain.Specification {
	var specifications []*domain.Specification

	for _, inputDto := range specification {
		spec := &domain.Specification{
			Name:        inputDto.Name,
			Description: inputDto.Description,
			CarID:       inputDto.CarID,
		}

		specifications = append(specifications, spec)
	}

	return specifications
}

func ConvertInputSpecificationToDomainUpdate(specification []*specificationdtos.SpecificationInputDto) []*domain.Specification {
	var specifications []*domain.Specification

	for _, inputDto := range specification {
		spec := &domain.Specification{
			Name:        inputDto.Name,
			Description: inputDto.Description,
			CarID:       inputDto.CarID,
		}

		specifications = append(specifications, spec)
	}

	return specifications
}

func ConvertSpecificationToDomainCreate(specification []*specificationdtos.SpecificationInputDto, id string) []*domain.Specification {
	var specifications []*domain.Specification

	for _, inputDto := range specification {
		spec := &domain.Specification{
			ID:          xid.New().String(),
			Name:        inputDto.Name,
			Description: inputDto.Description,
			CarID:       id,
		}

		specifications = append(specifications, spec)
	}

	return specifications
}
