package maintenanceutils

import (
	"errors"
	"sync"
	"time"
)

func ValidateDateRange(wg *sync.WaitGroup, errorChan chan<- error, startDate, endDate string) {
	defer wg.Done()

	startDateTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		errorChan <- err
		return
	}

	endDateTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		errorChan <- err
		return
	}

	if startDateTime.After(endDateTime) {
		errorChan <- errors.New("start date should be before or equal to end date")
	}
}
