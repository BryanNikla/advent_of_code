package utils

// AbsoluteValue - Helper function to calculate the absolute value
func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
