package usecases

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	repoutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/repo-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/rs/xid"
)

type PostCarUseCase struct {
	carRepository           repositories.CarRepository
	specificationRepository repositories.SpecificationRepository
}

func NewPostCarUseCase(
	carRepository repositories.CarRepository,
	specificationRepository repositories.SpecificationRepository) *PostCarUseCase {
	return &PostCarUseCase{
		carRepository:     carRepository,
		specificationRepository: specificationRepository,
	}
}

func (useCase *PostCarUseCase) Execute(inputDTO dtos.CarInputDTO) (*dtos.CarOutputDTO, error){
	if err := validation_errors.ValidateStruct(inputDTO); err != nil {
		return nil, err
	}

	carID := xid.New().String();

	specifications := repoutils.ConvertSpecificationToDomainCreate(inputDTO.Specification,carID)

	for _, spec := range specifications {
    fmt.Printf("%+v\n", spec)
}

	newCar := &domain.Car{  
		ID:            carID,
		Name:          inputDTO.Name,
		Description:   inputDTO.Description,
		DailyRate:     inputDTO.DailyRate,
		Available:     inputDTO.Available,
		LicensePlate:  inputDTO.LicensePlate,
		FineAmount:    inputDTO.FineAmount,
		Brand:         inputDTO.Brand,
		CategoryID:    inputDTO.CategoryID,
	}

	_, err := useCase.carRepository.FindCarByLicensePlate(newCar.LicensePlate)
	if err != nil {
		return nil, err
	}

	if err := useCase.carRepository.RegisterCar(newCar); err != nil {
		return nil, fmt.Errorf("failed to create car record: %w", err)
	}

	if err := useCase.specificationRepository.PostMultipleSpecifications(specifications); err != nil{
		return nil, fmt.Errorf("failed to create specification record: %w", err)
	}
	
	outPut := &dtos.CarOutputDTO{
		ID:           newCar.ID,
		Name:         newCar.Name,
		Description:  newCar.Description,
		DailyRate:    newCar.DailyRate,
		Available:    newCar.Available,
		LicensePlate: newCar.LicensePlate,
		FineAmount:   newCar.FineAmount,
		Brand:        newCar.Brand,
		CategoryID:   newCar.CategoryID,
		Specification: repoutils.ConvertSpecificationToDTO(specifications),
	}

	return outPut, nil
}

