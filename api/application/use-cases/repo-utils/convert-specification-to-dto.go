package repoutils

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	s "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller/specification-dtos"
)

func ConvertSpecificationToDTO(specificaion []*domain.Specification) []s.SpecificationOutputDto {
	var SpecificationsDTO []s.SpecificationOutputDto
	for _, spec := range specificaion {
		SpecificationsDTO = append(SpecificationsDTO, s.SpecificationOutputDto{
			ID:          spec.ID,
			Name:        spec.Name,
			Description: spec.Description,
			CarID:       spec.CarID,
		})
	}
	return SpecificationsDTO
}

func FuncConvertSpecificationToDomain(specification []s.SpecificationOutputDto) []*domain.Specification {
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
