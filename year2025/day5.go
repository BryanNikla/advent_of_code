package year2025

import (
	"sort"

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

func day5part1(input string) int {
	lines := utils.GetLines(input)
	countGoodIngredients := 0
	var freshRanges []utils.Range
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
			freshRanges = append(freshRanges, utils.NewRangeFromString(line, "-"))
		}
	}
	return countGoodIngredients
}

func day5part2(input string) int {

	var ranges []utils.Range
	for _, line := range utils.GetLines(input) {
		if line == "" {
			break
		}
		ranges = append(ranges, utils.NewRangeFromString(line, "-"))
	}

	var mergedRanges []utils.Range

	// Sorting the slices in order by start
	// Allows merging later & handling merging in one pass
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	if len(ranges) > 0 {
		mergedRanges = append(mergedRanges, ranges[0])
	}

	utils.ForEach(ranges, func(r utils.Range, _ int) {
		lastMergedIndex, lastMerged := utils.Last(mergedRanges)
		merged, didMerge := lastMerged.Merge(r)
		if didMerge {
			mergedRanges[lastMergedIndex] = merged
		} else {
			mergedRanges = append(mergedRanges, r)
		}
	})

	return utils.Reduce(mergedRanges, func(total int, r utils.Range, _ int) int {
		return total + r.Length()
	})
}
