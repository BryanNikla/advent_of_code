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

// ReplaceValues - Replace all instances of old value in slice with a new value
func ReplaceValues[V string | int](slice []V, old V, new V) []V {
	for i := range slice {
		if slice[i] == old {
			slice[i] = new
		}
	}
	return slice
}

// FindFirstSpanWithLength - Find the first span of indexes of a given slice that range a specified length for a target value
func FindFirstSpanWithLength[V string | int](nums []V, target V, length int) (int, int) {
	start := -1
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if count == 0 {
				start = i
			}
			count++
			if count == length {
				return start, i // Found a span with the given length
			}
		} else {
			count = 0 // Reset count if the span is broken
			start = -1
		}
	}
	return -1, -1 // No span found
}
