package year2025

import (
	"advent_of_code/utils"
)

func SolutionDay4() utils.Solution {
	input := getInput(4)
	return utils.Solution{
		Day:   4,
		Test1: day4part1(input) == 13,
		Test2: day4part2(input) == 43,
	}
}

const (
	PaperRoll           = "@"
	EmptySpace          = "."
	MaxAdjacentToAccess = 3
)

func day4part1(input string) int {
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))

	var total int
	utils.EachMatrix(matrix, func(obj string, coords utils.Coordinates, _ [][]string) {
		if obj == PaperRoll {
			var adjacent int
			utils.EachSurroundingInMatrix(matrix, coords, func(adjacentChar string, _ utils.Coordinates, _ [][]string) {
				if adjacentChar == PaperRoll {
					adjacent++
				}
			})
			if adjacent <= MaxAdjacentToAccess {
				total++
			}
		}
	})

	return total
}

func day4part2(input string) int {
	var total int
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))

	process := true
	for process { // Loop as long as we are replacing any paper rolls
		process = false
		utils.EachMatrix(matrix, func(obj string, coords utils.Coordinates, _ [][]string) {
			if obj == PaperRoll {
				var adjacent int
				utils.EachSurroundingInMatrix(matrix, coords, func(adjacentObj string, _ utils.Coordinates, _ [][]string) {
					if adjacentObj == PaperRoll {
						adjacent++
					}
				})
				if adjacent <= MaxAdjacentToAccess {
					total++
					process = true
					utils.SetAtMatrixPosition(matrix, coords, EmptySpace)
				}
			}

		})
	}

	return total
}
