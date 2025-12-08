package utils

import (
	"strconv"
)

// StringToInteger - Simpler string to integer that handles error (really just used to clean up solution logic)
func StringToInteger(input string) int {
	// Handle empty strings gracefully by returning 0
	if input == "" {
		return 0
	}
	integer, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return integer
}

func StringToFloat64(input string) float64 {
	// Handle empty strings gracefully by returning 0
	if input == "" {
		return 0
	}
	floatValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	return floatValue
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
