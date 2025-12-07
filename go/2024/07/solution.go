package year2024day07

import (
	"fmt"
	"strings"

	utils "advent_of_code/utils"
)

func Solve() utils.Solution {
	input1, input2 := utils.GetInputs(2024, 7)
	return utils.Solution{
		Day:   7,
		Test1: part1(input1) == 3749,
		Test2: part2(input2) == 11387,
	}
}

func part1(input string) int {
	var sum int
	for calibration, numbers := range prepareCalibrations(utils.GetLines(input)) {
		allOperatorCombinations := utils.AllPossibleCombinations(len(numbers)-1, []string{"+", "*"})
		for _, operators := range allOperatorCombinations {
			if calibration == executeOperationsOnNumbers(numbers, operators) {
				sum = sum + calibration
				break
			}
		}
	}
	return sum
}

func part2(input string) int {
	var sum int
	for calibration, numbers := range prepareCalibrations(utils.GetLines(input)) {
		allOperatorCombinations := utils.AllPossibleCombinations(len(numbers)-1, []string{"+", "*", "||"})
		for _, operators := range allOperatorCombinations {
			if calibration == executeOperationsOnNumbers(numbers, operators) {
				sum = sum + calibration
				break
			}
		}
	}
	return sum
}

func prepareCalibrations(lines []string) map[int][]int {
	calibrations := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		key := utils.StringToInteger(parts[0])
		numbersAsStrings := strings.Split(parts[1], " ")
		calibrations[key] = make([]int, 0, len(numbersAsStrings))
		for _, n := range numbersAsStrings {
			calibrations[key] = append(calibrations[key], utils.StringToInteger(n))
		}
	}
	return calibrations
}

func executeOperationsOnNumbers(numbers []int, operators []string) int {
	var i int
	for ii, number := range numbers {
		if ii == 0 {
			i = number
		} else {
			switch operators[ii-1] {
			case "+":
				i = i + number
			case "*":
				i = i * number
			case "||":
				i = utils.StringToInteger(fmt.Sprintf("%d%d", i, number))
			}
		}
	}
	return i
}
