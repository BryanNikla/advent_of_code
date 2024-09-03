package year2023

import (
	utilities "advent_of_code/common"
	"strconv"
	"strings"
)

func Day3() (int, int, bool, bool) {
	input := utilities.GetInputContent(2023, 3)
	test1 := utilities.GetTestContent(2023, 3, 1)
	// test2 := utilities.GetTestContent(2023, 3, 2)

	return part1(input), 0, part1(test1) == 4361, false
}

func part1(input string) int {
	var lines = utilities.GetLines(input)

	// make an empty matrix
	var matrix [][]string

	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	var parts []int

	var currentNumberString string = ""
	var currentNumberIsPart bool = false

	utilities.EachMatrix(matrix, func(char string, x int, y int, matrix [][]string) {

		if isDigit(char) {
			currentNumberString = currentNumberString + char

			// Only bother with this if we haven't already determined this is a part
			if !currentNumberIsPart {
				utilities.EachSurroundingInMatrix(matrix, x, y, func(c string, xx int, yy int, m [][]string) {
					if isSymbol(c) {
						currentNumberIsPart = true
					}
				})
			}
		}

		// If we're at the end of the number, or the end of the row add to parts if it's a part
		if utilities.IsLastColOfMatrix(matrix, x, y) || !isDigit(char) {
			if currentNumberIsPart {
				i, _ := strconv.Atoi(currentNumberString)
				parts = append(parts, i)
			}
			currentNumberIsPart = false
			currentNumberString = ""
		}
	})

	return utilities.SumValuesInSlice(parts)
}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func isSymbol(char string) bool {
	return char != "." && !isDigit(char)
}
