package specificationusecases

import (
	"errors"
	"fmt"

	repo "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/specification-controller/specification-dtos"
	"github.com/rs/xid"
)

func PostSpecificationUseCase(registerSpecification *dtos.SpecificationInputDto, specficationRepository repo.SpecificationRepository) (*dtos.SpecificationOutputDto, error) {
	existingSpecification, err := specficationRepository.FindSpecificationByName(registerSpecification.Name)
	if err != nil {
			return nil, fmt.Errorf("error querying specification: %w", err)
	}
	if existingSpecification != nil {
			return nil, errors.New("specification already exists")
	}
	
	newSpecification := &domain.Specification{
		ID:          xid.New().String(),
		Name:        registerSpecification.Name,
		Description: registerSpecification.Description,
		CarID:       registerSpecification.CarID,
	}

	if err := specficationRepository.PostSpecification(newSpecification); err != nil {
		return nil, fmt.Errorf("failed to create specification: %w", err)
	}

	outputDTO := &dtos.SpecificationOutputDto{
		ID:          newSpecification.ID,
		Name:        newSpecification.Name,
		Description: newSpecification.Description,
		CarID:       newSpecification.CarID,
	}
	return outputDTO, nil
}
