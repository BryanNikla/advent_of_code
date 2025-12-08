package utils

// Number defines a type constraint that includes all integer & floating-point types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

// AbsoluteValue - Helper function to calculate the absolute value
func AbsoluteValue[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func PositiveMod(val int, n int) int {
	return ((val % n) + n) % n
}

// Multiply takes any number of arguments of type T (where T is a Number)
// and returns the product as type T.
func Multiply[T Number](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}

	// Cast the untyped constant '1' to type T
	var total T = 1
	for _, n := range nums {
		total *= n
	}
	return total
}
