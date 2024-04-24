package maintenanceutils

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDateRange(t *testing.T) {
	tests := []struct {
		name      string
		startDate string
		endDate   string
		wantErr   error
	}{
		{
			name:      "Valid date range",
			startDate: "2024-01-01",
			endDate:   "2024-01-31",
			wantErr:   nil,
		},
		{
			name:      "Start date after end date",
			startDate: "2024-02-01",
			endDate:   "2024-01-31",
			wantErr:   errors.New("start date should be before or equal to end date"),
		},
		{
			name:      "Invalid start date format",
			startDate: "01-01-2024",
			endDate:   "2024-01-31",
			wantErr:   errors.New("parsing time \"01-01-2024\" as \"2006-01-02\": cannot parse \"01-01-2024\" as \"2006\""),
		},
		{
			name:      "Invalid end date format",
			startDate: "2024-01-01",
			endDate:   "31-01-2024",
			wantErr:   errors.New("parsing time \"31-01-2024\" as \"2006-01-02\": cannot parse \"31-01-2024\" as \"2006\""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorChan := make(chan error)
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go ValidateDateRange(wg, errorChan, tt.startDate, tt.endDate)

			wg.Wait()

			select {
			case err := <-errorChan:
				if err != nil {
					assert.Equal(t, tt.wantErr.Error(), err.Error())
				}
			default:
			}
		})
	}

}
