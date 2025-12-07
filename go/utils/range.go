package utils

import (
	"strings"
)

type Range struct {
	Start     int
	End       int
	Inclusive bool
}

func NewRange(start int, end int, inclusive bool) Range {
	return Range{
		Start:     start,
		End:       end,
		Inclusive: inclusive,
	}
}

// NewRangeFromString creates a new Range from a string representation like "5-10"
// Convenient as many ranges for AoC inputs tend to be represented this way.
func NewRangeFromString(input string, separator string) Range {
	startStr, endStr, _ := strings.Cut(input, separator)
	start := StringToInteger(startStr)
	end := StringToInteger(endStr)
	return NewRange(start, end, true)
}

func (r Range) Contains(value int) bool {
	if r.Inclusive {
		return value >= r.Start && value <= r.End
	}
	return value >= r.Start && value < r.End
}

func (r Range) ToSlice() []int {
	l := r.Length()
	if l <= 0 {
		return []int{} // Handle empty/invalid ranges
	}
	result := make([]int, 0, l)
	for i := r.Start; ; i++ {
		if r.Inclusive {
			if i > r.End {
				break
			}
		} else {
			if i >= r.End {
				break
			}
		}
		result = append(result, i)
		// Safety check for MaxInt to prevent infinite loop on overflow
		if i == int(^uint(0)>>1) {
			break
		}
	}
	return result
}

func (r Range) Length() int {
	var length int
	if r.Inclusive {
		length = r.End - r.Start + 1
	} else {
		length = r.End - r.Start
	}
	if length < 0 {
		return 0
	}
	return length
}

// Merge attempts to merge two ranges.
// It returns the merged range (normalized to Inclusive) and true if they overlap or are adjacent.
// Otherwise, it returns an empty Range and false.
func (r Range) Merge(other Range) (Range, bool) {
	// Normalize: Find the last included integer for both ranges
	rLast := r.End
	if !r.Inclusive {
		rLast--
	}
	otherLast := other.End
	if !other.Inclusive {
		otherLast--
	}

	// Check overlap/adjacency
	if r.Start > otherLast+1 || other.Start > rLast+1 {
		return Range{}, false
	}

	// Merge
	return Range{
		Start:     min(r.Start, other.Start),
		End:       max(rLast, otherLast),
		Inclusive: true,
	}, true
}
