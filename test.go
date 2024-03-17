package maintenanceusecases

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/dtos"
	"github.com/gin-gonic/gin"
)

type MaintenanceUseCaseFactory interface {
	CreateDeleteMaintenanceUseCase() DeleteMaintenanceUseCase
}

type DeleteMaintenanceUseCaseFactory struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewDeleteMaintenanceUseCaseFactory(
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) *DeleteMaintenanceUseCaseFactory {
	return &DeleteMaintenanceUseCaseFactory{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (factory *DeleteMaintenanceUseCaseFactory) CreateDeleteMaintenanceUseCase() DeleteMaintenanceUseCase {
	return NewDeleteMaintenanceUseCase(factory.carRepository, factory.maintenanceRepository)
}

type DeleteMaintenanceUseCase struct {
	carRepository         repositories.CarRepository
	maintenanceRepository repositories.MaintenanceRepository
}

func NewDeleteMaintenanceUseCase(
	carRepository repositories.CarRepository,
	maintenanceRepository repositories.MaintenanceRepository) DeleteMaintenanceUseCase {
	return DeleteMaintenanceUseCase{
		carRepository:         carRepository,
		maintenanceRepository: maintenanceRepository,
	}
}

func (useCase DeleteMaintenanceUseCase) Execute(maintenanceID string) error {
	return nil
}

func (useCase DeleteMaintenanceUseCase) performMaintenanceDeletion(errorChan chan<- error, resultChan chan<- *m.MaintenanceOutputDTO, maintenanceID string) {
}

// CONTROLLER
func DeleteMaintenanceController(context *gin.Context, factory MaintenanceUseCaseFactory) {
	maintenanceID := context.Param("maintenance_id")

	if maintenanceID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance ID"})
		return
	}

	deleteMaintenanceUseCase := factory.CreateDeleteMaintenanceUseCase()

	err := deleteMaintenanceUseCase.Execute(maintenanceID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, nil)
}
