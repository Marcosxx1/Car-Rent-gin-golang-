package utils

import (
	"testing"
)

func TestStringToFloat64_ValidInput(t *testing.T) {
	validCases := map[string]float64{
		"10.5": 10.5,
		"0":    0,
		"-3.7": -3.7,
	}

	for input, expected := range validCases {
		result, err := StringToFloat64(input)
		if err != nil {
			t.Errorf("Unexpected error for input %s: %v", input, err)
			continue
		}
		if result != expected {
			t.Errorf("Incorrect result for input %s. Expected: %f, Got: %f", input, expected, result)
		}
	}
}

func TestStringToFloat64_InvalidInput(t *testing.T) {
	invalidCases := []string{
		"abc",
		"10.5.2",
		"",
	}

	for _, input := range invalidCases {
		_, err := StringToFloat64(input)
		if err == nil {
			t.Errorf("Expected error for input %s", input)
		}
	}
}
