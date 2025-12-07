package year2024

import (
	"strconv"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2024, 2, func() utils.Solution {
		input1, input2 := utils.GetInput(2024, 2)
		return utils.Solution{
			Day:   2,
			Test1: day2part1(input1) == 2,
			Test2: day2part2(input2) == 4,
		}
	})
}

func day2part1(input string) int {
	var reports = linesToReports(utils.GetLines(input))
	return utils.Reduce(reports, func(safeReports int, report []int, _ int) int {
		if testReportAllSameDirection(report) && testReportSafeGaps(report) {
			return safeReports + 1
		}
		return safeReports
	})
}

func day2part2(input string) int {
	var reports = linesToReports(utils.GetLines(input))
	return utils.Reduce(reports, func(safeReports int, report []int, _ int) int {
		if recursiveTest(report, 0) {
			return safeReports + 1
		}
		return safeReports
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Recursively test for Part 2, testing that any iteration of this report with up to one item removed is safe
func recursiveTest(report []int, idx int) bool {
	adjustedReportToTest := removeFromReportAtIndex(report, idx)
	if testReportAllSameDirection(adjustedReportToTest) && testReportSafeGaps(adjustedReportToTest) {
		return true
	}
	if idx == len(report)-1 {
		return false
	}
	return recursiveTest(report, idx+1)
}

func linesToReports[report []int](lines []string) []report {
	reports := make([]report, 0, len(lines))
	for _, line := range lines {
		numbersAsStrings := strings.Split(line, " ")
		report := make(report, 0, len(numbersAsStrings))
		for _, num := range numbersAsStrings {
			number, _ := strconv.Atoi(num)
			report = append(report, number)
		}
		reports = append(reports, report)
	}
	return reports
}

func testReportAllSameDirection(report []int) bool {
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

func testReportSafeGaps(report []int) bool {
	for idx, number := range report {
		if idx > 0 {
			// Gap too large
			if utils.AbsoluteValue(number-report[idx-1]) > 3 {
				return false
			}
			// no gap
			if report[idx] == report[idx-1] {
				return false
			}
		}
	}
	return true
}

func removeFromReportAtIndex[report []int](r report, index int) report {
	var newReport = make(report, 0, len(r)-1)
	for i, x := range r {
		if i != index {
			newReport = append(newReport, x)
		}
	}
	return newReport
}
