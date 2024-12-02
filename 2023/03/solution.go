package year2023

import (
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func Day3() utils.Solution {
	input := utils.GetInputContent(2023, 3)
	// test2 := utilities.GetTestContent(2023, 3, 2)
	return utils.Solution{
		Part1: part1(input),
		Part2: 0,
		Test1: part1(utils.GetTestContent(2023, 3, 1)) == 4361,
	}
}

func part1(input string) int {
	var lines = utils.GetLines(input)

	// make an empty matrix
	var matrix [][]string

	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	var parts []int

	var currentNumberString string = ""
	var currentNumberIsPart bool = false

	utils.EachMatrix(matrix, func(char string, x int, y int, matrix [][]string) {

		if isDigit(char) {
			currentNumberString = currentNumberString + char

			// Only bother with this if we haven't already determined this is a part
			if !currentNumberIsPart {
				utils.EachSurroundingInMatrix(matrix, x, y, func(c string, xx int, yy int, m [][]string) {
					if isSymbol(c) {
						currentNumberIsPart = true
					}
				})
			}
		}

		// If we're at the end of the number, or the end of the row add to parts if it's a part
		if utils.IsLastColOfMatrix(matrix, x, y) || !isDigit(char) {
			if currentNumberIsPart {
				i, _ := strconv.Atoi(currentNumberString)
				parts = append(parts, i)
			}
			currentNumberIsPart = false
			currentNumberString = ""
		}
	})

	return utils.SumValuesInSlice(parts)
}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func isSymbol(char string) bool {
	return char != "." && !isDigit(char)
}
