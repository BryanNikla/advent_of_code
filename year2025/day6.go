package year2025

import (
	"sort"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func SolutionDay6() utils.Solution {
	input := getInput(6)
	return utils.Solution{
		Day:   6,
		Test1: day6part1(input) == 4277556,
		Test2: day6part2(input) == 3263827,
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func day6part1(input string) int {
	grandTotal := 0

	problems := make(map[int][]int)
	operations := make(map[int]string)

	for _, line := range utils.GetLines(input) {
		parts := strings.Fields(line)
		for i, part := range parts {
			if isNumeric(part) {
				problems[i] = append(problems[i], utils.StringToInteger(part))
			} else {
				operations[i] = part
			}
		}
	}

	for i := 0; i < len(problems); i++ {
		total := 0
		for j, num := range problems[i] {
			if j == 0 {
				total = num
			} else {
				switch operations[i] {
				case "+":
					total += num
				case "*":
					total *= num
				}
			}
		}
		grandTotal += total
	}

	return grandTotal
}

func day6part2(input string) int {
	grandTotal := 0
	lines := utils.GetLines(input)

	// Create a map of indices that are separators in the first line (initial candidates)
	indexIsSeparator := make(map[int]bool)
	for i, char := range strings.Split(lines[0], "") {
		if char == " " {
			indexIsSeparator[i] = true
		}
	}

	// Check all other lines to see if they also have " " at those indices
	for i := range indexIsSeparator {
		for _, line := range lines {
			lineChars := strings.Split(line, "")
			if lineChars[i] != " " {
				// This index is NOT a separator for all lines, remove from indexOfSeparators
				indexIsSeparator[i] = false
				break
			}
		}
	}

	// Collect all the indices that are separators for all lines
	splitIndices := []int{}
	for i, isSeparator := range indexIsSeparator {
		if isSeparator {
			splitIndices = append(splitIndices, i)
		}
	}

	// Sort the indices to ensure correct order
	sort.Ints(splitIndices)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////////

	problems := make(map[int][]string) // key is problem index, value is list of strings (to be processed later)
	operations := make(map[int]string) // key is problem index, value is operation

	// Now, split each line at those indices
	for lineIndex, line := range utils.GetLines(input) {
		var parts []string
		lastPosition := 0
		for _, i := range splitIndices {
			end := i
			if end > len(line) {
				end = len(line)
			}
			if lastPosition > end {
				lastPosition = end
			}
			parts = append(parts, line[lastPosition:end])
			lastPosition = end
		}
		if lastPosition < len(line) {
			parts = append(parts, line[lastPosition:])
		}

		for i, part := range parts {
			if lineIndex == len(lines)-1 {
				operations[i] = part
			} else {
				problems[i] = append(problems[i], part)
			}
		}
	}

	// Now lets process the "problems"
	for i := 0; i < len(problems); i++ {
		// Reconstruct the true strings by reading vertically
		trueStrings := []string{}
		for _, str := range problems[i] {
			digits := strings.Split(str, "")
			for di, digit := range digits {
				if di < len(trueStrings) {
					trueStrings[di] += digit
				} else {
					trueStrings = append(trueStrings, digit)
				}
			}
		}

		// Convert to actual numbers. Discarding any 0s
		trueNumbers := []int{}
		for _, str := range trueStrings {
			number := utils.StringToInteger(strings.TrimSpace(str))
			if number > 0 {
				trueNumbers = append(trueNumbers, number)
			}
		}

		// DO THE MATHS
		total := 0
		for j, num := range trueNumbers {
			if j == 0 {
				total = num
			} else {
				switch strings.TrimSpace(operations[i]) {
				case "+":
					total += num
				case "*":
					total *= num
				}
			}
		}
		grandTotal += total
	}

	return grandTotal
}
