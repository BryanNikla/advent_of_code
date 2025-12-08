package utils

// MiddleItem - Returns the middle item of a slice. If even number of items, returns the upper middle
func MiddleItem[V any](slice []V) V {
	middleIndex := len(slice) / 2
	return slice[middleIndex]
}

// Reduce - Reducer for Slices. Initial value is always default empty for value type
func Reduce[T any, V any](slice []T, reducer func(accumulated V, currentValue T, currentIndex int) V) V {
	var value V
	for i := 0; i < len(slice); i++ {
		value = reducer(value, slice[i], i)
	}
	return value
}

// ForEach - Iterates over each element in the slice, executing the callback function for each element
// Basically mimics Javascript's Array.forEach() method
func ForEach[T any](slice []T, callback func(T, int)) {
	for i, element := range slice {
		callback(element, i)
	}
}

// Every - Returns true if all elements in the slice satisfy the predicate, and false otherwise
// Basically mimics Javascript's Array.every() method
func Every[T any](slice []T, predicate func(T, int) bool) bool {
	for i, element := range slice {
		if !predicate(element, i) {
			return false
		}
	}
	return true
}

// SumValuesInSlice - Sums all numeric values in a slice and returns the total
func SumValuesInSlice[V int | int8 | int16 | int32 | int64 | float32 | float64, S []V](slice S) V {
	var sum V
	for _, value := range slice {
		sum += value
	}
	return sum
}

// ReplaceValues - Replace all instances of old value in slice with a new value
func ReplaceValues[V string | int](slice []V, old V, new V) []V {
	for i := range slice {
		if slice[i] == old {
			slice[i] = new
		}
	}
	return slice
}

// FindFirstSpanWithLength - Find the first span of indexes of slice that range a specified length for a target value
func FindFirstSpanWithLength[V string | int](slice []V, target V, length int) (int, int) {
	var startIndex int
	var count int
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			if count == 0 {
				startIndex = i
			}
			count++
			// Found a span with the required length
			if count == length {
				return startIndex, i
			}
		} else {
			// Span broken, reset our counter
			count = 0
		}
	}
	// No span found
	return -1, -1
}

// Last - Return the last index and value of a slice, or -1 & zero value if empty
func Last[V any](slice []V) (int, V) {
	if len(slice) == 0 {
		var zero V
		return -1, zero
	}
	idx := len(slice) - 1
	return idx, slice[idx]
}
