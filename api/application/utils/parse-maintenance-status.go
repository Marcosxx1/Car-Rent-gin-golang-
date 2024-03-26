package utils

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain/enums"
)


func ParseMaintenanceStatus(value string) (enums.MaintenanceStatus, error) {
	switch value {
	case string(enums.Scheduled):
		return enums.Scheduled, nil
	case string(enums.InProgress):
		return enums.InProgress, nil
	case string(enums.Completed):
		return enums.Completed, nil
	case string(enums.PendingApproval):
		return enums.PendingApproval, nil
	case string(enums.Canceled):
		return enums.Canceled, nil
	case string(enums.AwaitingParts):
		return enums.AwaitingParts, nil
	case string(enums.AwaitingPayment):
		return enums.AwaitingPayment, nil
	case string(enums.Rescheduled):
		return enums.Rescheduled, nil
	case string(enums.MaintenanceFailed):
		return enums.MaintenanceFailed, nil
	case string(enums.AwaitingInspection):
		return enums.AwaitingInspection, nil
	default:
		return "", fmt.Errorf("invalid maintenance status")
	}
}