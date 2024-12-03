package year2023

import (
	"regexp"
	"strconv"

	"advent_of_code/utils"
)

var testSolutionPart1 int = 142
var testSolutionPart2 int = 281

func Day1() utils.Solution {
	input := utils.GetInputContent(2023, 1)
	test1 := utils.GetTestContent(2023, 1, 1)
	test2 := utils.GetTestContent(2023, 1, 2)

	return utils.Solution{
		Day:   1,
		Part1: strconv.Itoa(part1(input)),
		Part2: strconv.Itoa(part2(input)),
		Test1: part1(test1) == testSolutionPart1,
		Test2: part2(test2) == testSolutionPart2,
	}
}

func part1(input string) int {
	lines := utils.GetLines(input)

	var numbers []string

	for _, line := range lines {
		match := regexp.MustCompile(`\d`).FindAllString(line, -1)
		var first = match[0]
		var last = match[len(match)-1]
		numbers = append(numbers, first+last)
	}

	return sumSliceOfDigitStrings(numbers)
}

func part2(input string) int {
	lines := utils.GetLines(input)

	var numbers []string

	for _, line := range lines {
		r := regexp.MustCompile(`^(one|two|three|four|five|six|seven|eight|nine|\d)`)
		var lineNumbers []string
		for len(line) > 0 {
			match := r.FindString(line)
			line = line[1:]
			if match != "" {
				lineNumbers = append(lineNumbers, findDigit(match))
			}
		}
		if len(lineNumbers) != 0 {
			var first = lineNumbers[0]
			var last = lineNumbers[len(lineNumbers)-1]
			numbers = append(numbers, first+last)
		}
	}

	return sumSliceOfDigitStrings(numbers)
}

// sumSliceOfDigitStrings sums the digits in a slice of strings, assuming each string represents a Number digit
func sumSliceOfDigitStrings(digits []string) int {
	var sum int
	for _, digit := range digits {
		i, _ := strconv.Atoi(digit)
		sum = sum + i
	}
	return sum
}

// findDigit returns the digit if the string is a number word, else return the string
func findDigit(x string) string {
	switch x {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return x
	}
}
