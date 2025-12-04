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
	PaperRoll  = "@"
	EmptySpace = "."
)

func day4part1(input string) int {
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))

	var total int
	utils.EachMatrix(matrix, func(char string, coords utils.Coordinates, _ [][]string) {
		if char == "@" {
			var adjacent int
			utils.EachSurroundingInMatrix(matrix, coords, func(adjacentChar string, _ utils.Coordinates, _ [][]string) {
				if adjacentChar == "@" {
					adjacent++
				}
			})

			if adjacent < 4 {
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

	var replaced = 1 // start at one to trigger loop

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
				if adjacent < 4 {
					total++
					replaced++
					utils.SetAtMatrixPosition(nextMatrix, coords, EmptySpace)
				}
			}

		})
	}

	fmt.Printf("ANSWER: %d", total)
	return total
}
