package year2023

import (
	"strconv"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2023, 3, func() utils.Solution {
		input1, _ := utils.GetInput(2023, 3)
		return utils.Solution{
			Day:   3,
			Test1: day3part1(input1) == 4361,
		}
	})
}

func day3part1(input string) int {
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))

	var parts []int

	var currentNumberString string
	var currentNumberIsPart bool

	utils.EachMatrix(matrix, func(char string, coordinates utils.Coordinates, matrix [][]string) {

		if isDigit(char) {
			currentNumberString = currentNumberString + char

			// Only bother with this if we haven't already determined this is a part
			if !currentNumberIsPart {
				utils.EachSurroundingInMatrix(matrix, coordinates, func(char string, _ utils.Coordinates, _ [][]string) {
					if isSymbol(char) {
						currentNumberIsPart = true
					}
				})
			}
		}

		// If we're at the end of the number, or the end of the row add to parts if it's a part
		if utils.IsLastColOfMatrix(matrix, coordinates) || !isDigit(char) {
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
