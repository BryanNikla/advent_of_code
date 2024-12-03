package year2024day01

import (
	"sort"
	"strings"

	"advent_of_code/utils"
)

func Solve() utils.Solution {
	input, testInput1, testInput2 := utils.GetAllInputs(2024, 1)
	return utils.Solution{
		Day:   1,
		Part1: part1(input),
		Part2: part2(input),
		Test1: part1(testInput1) == 11,
		Test2: part2(testInput2) == 31,
	}
}

func part1(input string) int {
	lines := utils.GetLines(input)
	leftNumbers, rightNumbers := linesToNumberSlices(lines)
	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var total int
	for i, left := range leftNumbers {
		total += utils.AbsoluteValue(left - rightNumbers[i])
	}
	return total
}

func part2(input string) int {
	lines := utils.GetLines(input)
	leftNumbers, rightNumbers := linesToNumberSlices(lines)

	// Create a map of value occurrences in the rightNumbers slice
	occurrences := make(map[int]int)
	for _, v := range rightNumbers {
		occurrences[v]++
	}

	// Sum of a "similarity scores" for this problem.
	// A similarity score is found by "adding up each number in the left list after multiplying
	// it by the number of times that number appears in the right list"
	var similarityScoreTotal int
	for _, number := range leftNumbers {
		similarityScoreTotal += number * occurrences[number]
	}
	return similarityScoreTotal
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Values used in both parts of this project
func linesToNumberSlices(lines []string) ([]int, []int) {
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, utils.StringToInteger(parts[0]))
		right = append(right, utils.StringToInteger(parts[1]))
	}
	return left, right
}
