package year2025

import (
	utils "advent_of_code/utils"
)

func SolutionDay1() utils.Solution {
	inputStr := getInput(1)
	return utils.Solution{
		Day:   1,
		Test1: day1part1(inputStr) == 3,
		Test2: day1part2(inputStr) == 6,
	}
}

const (
	DAY1_LEFT                = "L"
	DAY1_RIGHT               = "R"
	DAY1_DIAL_START_POSITION = 50
	DAY1_DIAL_DIGITS         = 100 // 0-99
)

func day1part1(input string) int {
	lines := utils.GetLines(input)
	var endPositionsAtZero int
	dialPosition := DAY1_DIAL_START_POSITION
	for _, line := range lines {
		direction, distance := lineToDirectionAndDistance(line)
		if direction == DAY1_RIGHT {
			dialPosition = utils.PositiveMod(dialPosition+distance, DAY1_DIAL_DIGITS)
		} else if direction == DAY1_LEFT {
			dialPosition = utils.PositiveMod(dialPosition-distance, DAY1_DIAL_DIGITS)
		}
		if dialPosition == 0 {
			endPositionsAtZero++
		}
	}
	return endPositionsAtZero
}

func day1part2(input string) int {
	lines := utils.GetLines(input)
	var totalVisitsToZero int
	dialPosition := DAY1_DIAL_START_POSITION
	for _, line := range lines {
		direction, distance := lineToDirectionAndDistance(line)
		if direction == DAY1_RIGHT {
			visits, newPos := moveAndCountZeros(dialPosition, distance, DAY1_DIAL_DIGITS, false)
			totalVisitsToZero += visits
			dialPosition = newPos
		} else if direction == DAY1_LEFT {
			visits, newPos := moveAndCountZeros(dialPosition, distance, DAY1_DIAL_DIGITS, true)
			totalVisitsToZero += visits
			dialPosition = newPos
		}
	}
	return totalVisitsToZero
}

func lineToDirectionAndDistance(line string) (string, int) {
	direction := line[:1]
	distance := utils.StringToInteger(line[1:])
	return direction, distance
}

func moveAndCountZeros(startPos int, distance int, n int, isMovingLeft bool) (int, int) {
	visits := distance / n
	remainder := distance % n
	var newPos int
	if isMovingLeft {
		newPos = utils.PositiveMod(startPos-remainder, n)
	} else {
		newPos = (startPos + remainder) % n
	}
	if isMovingLeft {
		if startPos > 0 && (startPos-remainder <= 0) {
			visits++
		}
	} else {
		if startPos+remainder >= n {
			visits++
		}
	}
	return visits, newPos
}
