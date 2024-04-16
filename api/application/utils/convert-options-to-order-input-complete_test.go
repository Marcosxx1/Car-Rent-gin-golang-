package utils

import (
	"net/url"
	"testing"
	"time"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
)

func TestConvertOptionsToOrderInputComplete(t *testing.T) {
	queryParams := url.Values{
		"ID":              {"123"},
		"userID":          {"456"},
		"carID":           {"789"},
		"rentalStartDate": {"2024-04-08T10:00:00Z"},
		"rentalEndDate":   {"2024-04-10T10:00:00Z"},
		"totalCost":       {"150.5"},
		"orderStatus":     {"true"},
	}

	result, err := ConvertOptionsToOrderInputComplete(queryParams)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := &orderdto.OrderOutputDTO{
		ID:              "123",
		UserID:          "456",
		CarID:           "789",
		RentalStartDate: time.Date(2024, 4, 8, 10, 0, 0, 0, time.UTC),
		RentalEndDate:   time.Date(2024, 4, 10, 10, 0, 0, 0, time.UTC),
		TotalCost:       150.5,
		OrderStatus:     true,
	}

	if *result != *expected {
		t.Errorf("Result does not match expected. Got: %+v, Expected: %+v", result, expected)
	}
}

func TestConvertOptionsToOrderInputComplete_ErrorCases(t *testing.T) {
	invalidDateParams := url.Values{
		"rentalStartDate": {"2024-04-08"},
		"rentalEndDate":   {"2024-04-10T10:00:00Z"},
	}

	_, err := ConvertOptionsToOrderInputComplete(invalidDateParams)
	if err == nil {
		t.Error("Expected error for invalid date format")
	}

	invalidTotalCostParams := url.Values{
		"totalCost": {"invalid_cost"},
	}

	_, err = ConvertOptionsToOrderInputComplete(invalidTotalCostParams)
	if err == nil {
		t.Error("Expected error for invalid total cost format")
	}

	invalidOrderStatusParams := url.Values{
		"orderStatus": {"invalid_status"},
	}

	_, err = ConvertOptionsToOrderInputComplete(invalidOrderStatusParams)
	if err == nil {
		t.Error("Expected error for invalid order status format")
	}
}
