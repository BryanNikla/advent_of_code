package utils

func MiddleItem[V any](slice []V) V {
	middleIndex := len(slice) / 2
	return slice[middleIndex]
}

// Reduce - Reducer for Slices. Initial value is always default for value type
func Reduce[T any, V any](slice []T, reducer func(accumulated V, currentValue T, currentIndex int) V) V {
	var value V
	for i := 0; i < len(slice); i++ {
		value = reducer(value, slice[i], i)
	}
	return value
}

// Every - Returns true if all elements in the slice satisfy the predicate, and false otherwise
// Basically mimics Javascript's Array.every() method
func Every[T any](slice []T, predicate func(T) bool) bool {
	for _, element := range slice {
		if !predicate(element) {
			return false
		}
	}
	return true
}

func SumValuesInSlice[V int | int8 | int16 | int32 | int64 | float32 | float64, S []V](slice S) V {
	var sum V
	for _, value := range slice {
		sum += value
	}
	return sum
}
