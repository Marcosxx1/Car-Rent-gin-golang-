package utils

import "strconv"

func StringToFloat64(numberAsString string) (float64, error) {
	parsedFloat, err := strconv.ParseFloat(numberAsString, 64)
	if err != nil {
		return 0, err
	}
	return parsedFloat, nil
}
