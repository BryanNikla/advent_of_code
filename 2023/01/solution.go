package year2023

import (
	utilities "advent_of_code/common"
	"regexp"
	"strconv"
)

var testSolutionPart1 int = 142
var testSolutionPart2 int = 281

func Day1() (string, string, bool, bool) {
	input := utilities.GetInputContent(2023, 1)
	test1 := utilities.GetTestContent(2023, 1, 1)
	test2 := utilities.GetTestContent(2023, 1, 2)

	return strconv.Itoa(part1(input)), strconv.Itoa(part2(input)), part1(test1) == testSolutionPart1, part2(test2) == testSolutionPart2
}

func part1(input string) int {
	lines := utilities.GetLines(input)

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
	lines := utilities.GetLines(input)

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
