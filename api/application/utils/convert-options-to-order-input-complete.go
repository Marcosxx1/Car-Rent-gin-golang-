package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
)

// ConvertOptionsToOrderInputComplete converts URL query parameters to an OrderOutputDTO.
//
// It extracts values from the provided URL query parameters and populates an OrderOutputDTO.
// The function returns the populated OrderOutputDTO and an error if any parsing fails.
func ConvertOptionsToOrderInputComplete(queryParams url.Values) (*orderdto.OrderOutputDTO, error) {
	// Initialize searchOptions with default values
	searchOptions := orderdto.OrderOutputDTO{}

	// Extract values from queryParams
	for key, value := range queryParams {
		switch key {
		case "ID":
			searchOptions.ID = value[0]
		case "userID":
			searchOptions.UserID = value[0]
		case "carID":
			searchOptions.CarID = value[0]
		case "rentalStartDate":
			startDate, err := time.Parse(time.RFC3339, value[0])
			if err != nil {
				return &searchOptions, fmt.Errorf("invalid rental start date format: %v", err)
			}
			searchOptions.RentalStartDate = startDate
		case "rentalEndDate":
			endDate, err := time.Parse(time.RFC3339, value[0])
			if err != nil {
				return &searchOptions, fmt.Errorf("invalid rental end date format: %v", err)
			}
			searchOptions.RentalEndDate = endDate
		case "totalCost":
			totalCost, err := strconv.ParseFloat(value[0], 64)
			if err != nil {
				return &searchOptions, fmt.Errorf("invalid total cost format: %v", err)
			}
			searchOptions.TotalCost = totalCost
		case "orderStatus":
			orderStatus, err := strconv.ParseBool(value[0])
			if err != nil {
				return &searchOptions, fmt.Errorf("invalid order status format: %v", err)
			}
			searchOptions.OrderStatus = orderStatus
		}
	}

	return &searchOptions, nil
}
