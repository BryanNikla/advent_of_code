package go2024

import (
	"sort"
	"strings"

	utilities "advent_of_code/common"
)

var testSolution1 = 11
var testSolution2 = 31

func Day1() (any, any, bool, bool) {
	input := utilities.GetInputContent(2024, 1)
	test1 := utilities.GetTestContent(2024, 1, 1)
	test2 := utilities.GetTestContent(2024, 1, 1) // same test data for each part

	return part1(input),
		part2(input),
		part1(test1) == testSolution1,
		part2(test2) == testSolution2
}

func part1(input string) int {
	lines := utilities.GetLines(input)
	leftNumbers, rightNumbers := linesToNumberSlices(lines)
	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var total int
	for i, left := range leftNumbers {
		total += utilities.AbsoluteValue(left - rightNumbers[i])
	}
	return total
}

func part2(input string) int {
	lines := utilities.GetLines(input)
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

// Values used in both parts of this project
func linesToNumberSlices(lines []string) ([]int, []int) {
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, utilities.StringToInteger(parts[0]))
		right = append(right, utilities.StringToInteger(parts[1]))
	}
	return left, right
}
