package specificationusecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller/specification-dtos"
	"github.com/rs/xid"
)

type PostSpecificationUseCase struct {
	specificationRepository repositories.SpecificationRepository
}

func NewPostSpecificationUseCase(specificationRepository repositories.SpecificationRepository) *PostSpecificationUseCase {
	return &PostSpecificationUseCase{
		specificationRepository: specificationRepository,
	}
}

func (useCase *PostSpecificationUseCase) Execute(registerSpecification *dtos.SpecificationInputDto) (*dtos.SpecificationOutputDto, error) {

	existingSpecification, err := useCase.specificationRepository.FindSpecificationByName(registerSpecification.Name)
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

	if err := useCase.specificationRepository.PostSpecification(newSpecification); err != nil {
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
