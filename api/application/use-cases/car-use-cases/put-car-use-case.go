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



	carToBeUpdated := &domain.Car{
		ID:            id,
		Name:          updateRequest.Name,
		Description:   updateRequest.Description,
		DailyRate:     updateRequest.DailyRate,
		Available:     updateRequest.Available,
		LicensePlate:  updateRequest.LicensePlate,
		FineAmount:    updateRequest.FineAmount,
		Brand:         updateRequest.Brand,
		Specification: repoutils.FuncConvertSpecificationToDomain(updateRequest.Specification),
	}

	if err := validation_errors.ValidateStruct(carToBeUpdated); err != nil {
		return nil, err
	}

	car, err := useCase.carRepository.UpdateCar(id, *carToBeUpdated)
	if err != nil {
		return nil, err
	}

	carToBeReturned := &dtos.CarOutputDTO{
		ID:           car.ID,
		Name:         car.Name,
		Description:  car.Description,
		Available:    car.Available,
		LicensePlate: car.LicensePlate,
		FineAmount:   car.FineAmount,
		Brand:        car.Brand,
		Specification: repoutils.ConvertSpecificationToDTO(car.Specification),

	}

	/* 	println(carToBeReturned)
	   	fmt.Printf("%+v\n", carToBeReturned) */

	return carToBeReturned, nil
}
