package utils

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

func (r Range) Contains(value int) bool {
	if r.Inclusive {
		return value >= r.Start && value <= r.End
	}
	return value >= r.Start && value < r.End
}

func (r Range) ToSlice() []int {
	var result []int
	end := r.End
	if r.Inclusive {
		end++
	}
	for i := r.Start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

func (r Range) Length() int {
	if r.Inclusive {
		return r.End - r.Start + 1
	}
	return r.End - r.Start
}
