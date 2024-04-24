package maintenanceutils

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain/enums"
)

func TestPerformMaintenanceValidation(t *testing.T) {

	tests := []struct {
		name                string
		inputDTO            maintenancedtos.MaintenanceInputDTO
		wantValidationError bool
	}{
		{
			name: "Valid input",
			inputDTO: maintenancedtos.MaintenanceInputDTO{
				CarID:                     "1",
				MaintenanceType:           "Oil Change",
				OdometerReading:           10000,
				ScheduledMaintenance:      true,
				MaintenanceStatus:         enums.Completed,
				LastMaintenanceDate:       time.Now().AddDate(0, 0, -30),
				NextMaintenanceDueDate:    time.Now().AddDate(0, 1, 0),
				MaintenanceCompletionDate: time.Now(),
				Parts: []maintenancedtos.PartInputDTO{
					{
						Name:            "Oil Filter",
						Cost:            20,
						Quantity:        1,
						ReplacementDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantValidationError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			errorChan := make(chan error)
			validationErrorSignal := make(chan bool)

			wg.Add(1)
			go PerformMaintenanceValidation(wg, errorChan, validationErrorSignal, tt.inputDTO)

			wg.Wait()

			select {
			case err := <-errorChan:
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantValidationError, <-validationErrorSignal)
			default:
				assert.False(t, tt.wantValidationError, "Expected no validation error")
			}
		})
	}
}
