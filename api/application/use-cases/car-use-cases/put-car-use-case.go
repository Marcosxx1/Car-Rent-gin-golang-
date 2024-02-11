package usecases

import (
	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

//updateCarUseCase := *usecases.NewUpdateCarUseCase(carRepository, specificationRepository)

type PutCarUseCase struct {
	carRepository           r.CarRepository
	specificationRepository r.SpecificationRepository
}

func NewUpdateCarUseCase(
	carRepository r.CarRepository,
	specificationRepository r.SpecificationRepository) *PutCarUseCase {

	return &PutCarUseCase{
		carRepository:           carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *PutCarUseCase) Execute(id string, updateRequest *dtos.CarOutputDTO) (*dtos.CarOutputDTO, error) {

	if err := validation_errors.ValidateStruct(updateRequest); err != nil {
		return nil, err
	}

	var domainSpecification []*domain.Specification

	if len(updateRequest.Specification) > 0 {
		domainSpecification = repoutils.ConvertSpecificationToDomainUpdate(updateRequest.Specification)

		if err := validation_errors.ValidateStruct(domainSpecification); err != nil {
			return nil, err
		}
	}

	carToBeUpdated := &domain.Car{
		ID:           id,
		Name:         updateRequest.Name,
		Description:  updateRequest.Description,
		DailyRate:    updateRequest.DailyRate,
		Available:    updateRequest.Available,
		LicensePlate: updateRequest.LicensePlate,
		FineAmount:   updateRequest.FineAmount,
		Brand:        updateRequest.Brand,
	}

	carUpdated, err := useCase.carRepository.UpdateCar(id, *carToBeUpdated)
	if err != nil {
		return nil, err
	}

	specificationUpdated, err := useCase.specificationRepository.UpdateSpecification(id, domainSpecification)
	if err != nil {
		return nil, err
	}

	carToBeReturned := &dtos.CarOutputDTO{
		ID:            carUpdated.ID,
		Name:          carUpdated.Name,
		Description:   carUpdated.Description,
		Available:     carUpdated.Available,
		LicensePlate:  carUpdated.LicensePlate,
		FineAmount:    carUpdated.FineAmount,
		Brand:         carUpdated.Brand,
		Specification: repoutils.ConvertSpecificationToDTO(specificationUpdated),
	}

	return carToBeReturned, nil
}
