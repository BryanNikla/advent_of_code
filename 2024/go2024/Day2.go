package go2024

import (
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func Day2() utils.Solution {
	input := utils.GetInputContent(2024, 2)
	testInput := utils.GetTestContent(2024, 2, 1) // Test 1 & 2 use same data

	return utils.Solution{
		Part1: day2part1(input),
		Part2: day2part2(input),
		Test1: day2part1(testInput) == 2,
		Test2: day2part2(testInput) == 4,
	}
}

func day2part1(input string) int {
	lines := utils.GetLines(input)
	reports := day2_linesToReports(lines)
	var safeReports int
	for _, report := range reports {
		var isUnsafe bool
		var isGoingDown bool
		for idx, number := range report {
			if idx != 0 {
				if isGoingDown {
					if number > report[idx-1] {
						isUnsafe = true
					}
				} else {
					if number < report[idx-1] {
						isUnsafe = true
					}
				}

				if utils.AbsoluteValue(number-report[idx-1]) > 3 {
					isUnsafe = true
				}
				if report[idx] == report[idx-1] {
					isUnsafe = true
				}
			} else {
				if number > report[idx+1] {
					isGoingDown = true
				}
			}
		}

		if !isUnsafe {
			safeReports++
		}
	}
	return safeReports
}

func day2part2(input string) int {
	lines := utils.GetLines(input)
	reports := day2_linesToReports(lines)
	var safeReports int
	for _, report := range reports {

	}

	return safeReports
}

func day2_linesToReports(lines []string) [][]int {
	reports := make([][]int, 0, len(lines))
	for _, line := range lines {
		numbersAsStrings := strings.Split(line, " ")
		report := make([]int, 0, len(numbersAsStrings))
		for _, num := range numbersAsStrings {
			number, _ := strconv.Atoi(num)
			report = append(report, number)
		}
		reports = append(reports, report)
	}
	return reports
}

func day2_testReport_allSameDirection(report []int) bool {
	var isGoingDown bool
	for idx, number := range report {
		if idx == 0 {
			if number > report[idx+1] {
				isGoingDown = true
			}
		} else {
			if isGoingDown {
				if number > report[idx-1] {
					return false
				}
			} else {
				if number < report[idx-1] {
					return false
				}
			}
		}
	}
	return true
}
