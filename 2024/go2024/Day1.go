package go2024

import (
	"sort"
	"strings"

	"advent_of_code/utils"
)

var testSolution1 = 11
var testSolution2 = 31

func Day1() utils.Solution {
	input := utils.GetInputContent(2024, 1)
	testInput := utils.GetTestContent(2024, 1, 1) // Test 1 & 2 use same data

	return utils.Solution{
		Part1: day1part1(input),
		Part2: day1part2(input),
		Test1: day1part1(testInput) == 11,
		Test2: day1part2(testInput) == 31,
	}
}

func day1part1(input string) int {
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

func day1part2(input string) int {
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
