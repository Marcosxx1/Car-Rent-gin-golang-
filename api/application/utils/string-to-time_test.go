package utils

import (
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	tests := []struct {
		description string
		dateString  string
		expected    time.Time
		expectError bool
	}{
		{
			description: "Valid date string",
			dateString:  "2024-04-08 12:00:00",
			expected:    time.Date(2024, 4, 8, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			description: "Invalid date string",
			dateString:  "invalid_date",
			expected:    time.Time{},
			expectError: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			result, err := StringToTime(testCase.dateString)

			if (err != nil) != testCase.expectError {
				t.Errorf("StringToTime() error = %v, expected error = %v", err, testCase.expectError)
				return
			}

			if !result.Equal(testCase.expected) {
				t.Errorf("StringToTime() result = %v, expected = %v", result, testCase.expected)
			}
		})
	}
}
