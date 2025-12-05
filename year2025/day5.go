package year2025

import (
	"fmt"
	"sort"
	"strings"

	"advent_of_code/utils"
)

func SolutionDay5() utils.Solution {
	input := getInput(5)
	return utils.Solution{
		Day:   5,
		Test1: day5part1(input) == 3,
		Test2: day5part2(input) == 14,
	}
}

type FreshIdRange struct {
	Start int
	End   int
}

func (r FreshIdRange) Merge(other FreshIdRange) (FreshIdRange, bool) {
	if r.Start > other.End+1 || other.Start > r.End+1 {
		return FreshIdRange{}, false
	}
	newRange := FreshIdRange{
		Start: min(r.Start, other.Start),
		End:   max(r.End, other.End),
	}
	return newRange, true
}

func day5part1(input string) int {
	lines := utils.GetLines(input)
	countGoodIngredients := 0
	var freshRanges []FreshIdRange
	var parseIngredients bool
	for _, line := range lines {
		if line == "" {
			parseIngredients = true
			continue
		}
		if parseIngredients {
			ingredientID := utils.StringToInteger(line)
			for _, r := range freshRanges {
				if utils.InRange(ingredientID, r.Start, r.End) {
					countGoodIngredients++
					break
				}
			}
		} else {
			bounds := strings.Split(line, "-")
			start := utils.StringToInteger(bounds[0])
			end := utils.StringToInteger(bounds[1])
			freshRanges = append(freshRanges, FreshIdRange{Start: start, End: end})
		}
	}
	return countGoodIngredients
}

func day5part2(input string) int {

	var ranges []FreshIdRange
	for _, line := range utils.GetLines(input) {
		if line == "" {
			break
		}
		bounds := strings.Split(line, "-")
		ranges = append(ranges, FreshIdRange{
			Start: utils.StringToInteger(bounds[0]),
			End:   utils.StringToInteger(bounds[1]),
		})
	}

	var mergedRanges []FreshIdRange

	// Sorting the slices in order by start
	// Allows merging later & handling merging in one pass
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	if len(ranges) > 0 {
		mergedRanges = append(mergedRanges, ranges[0])
	}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		lastMergedIndex, lastMerged := utils.Last(mergedRanges)

		merged, didMerge := lastMerged.Merge(current)
		if didMerge {
			mergedRanges[lastMergedIndex] = merged
		} else {
			mergedRanges = append(mergedRanges, current)
		}
	}

	totalFreshIngredients := 0
	for _, r := range mergedRanges {
		totalFreshIngredients += r.End - r.Start + 1 // Adding +1 to be inclusive of ends
	}

	fmt.Printf("\nTotal Fresh Ingredients: %d\n", totalFreshIngredients)
	return totalFreshIngredients
}
