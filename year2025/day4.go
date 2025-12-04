package year2025

import (
	"fmt"

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

	fmt.Println(total)
	return total
}

func day4part2(input string) int {
	var total int
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))

	nextMatrix := utils.CloneMatrix(matrix)

	var replaced = 1 // Initialize as 1 to trigger first pass

	// Loop as long as we are replacing any paper rolls
	// Stop when nothing was removed in a full pass
	for replaced > 0 {
		replaced = 0
		matrix = nextMatrix

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
					replaced++
					utils.SetAtMatrixPosition(nextMatrix, coords, EmptySpace)
				}
			}

		})
	}

	return total
}
