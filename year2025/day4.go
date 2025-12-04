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

type WarehouseFloor [][]string

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
			if isPaperRollAccessible(matrix, coords) {
				total++
			}
		}
	})
	return total
}

func day4part2(input string) int {
	var total int
	room := WarehouseFloor(utils.LinesToCharacterMatrix(utils.GetLines(input)))

	var processQueue []utils.Coordinates

	utils.EachMatrix(room, func(obj string, coords utils.Coordinates, _ WarehouseFloor) {
		if obj == PaperRoll {
			if isPaperRollAccessible(room, coords) {
				processQueue = append(processQueue, coords)
			}
		}
	})

	for len(processQueue) > 0 {
		curr := processQueue[0]
		processQueue = processQueue[1:]

		if utils.GetValueAtCords(room, curr) == PaperRoll {
			removePaperRoll(room, curr)
			total++
			utils.EachSurroundingInMatrix(room, curr, func(adjacentObj string, adjacentCoords utils.Coordinates, _ WarehouseFloor) {
				if adjacentObj == PaperRoll {
					if isPaperRollAccessible(room, adjacentCoords) {
						processQueue = append(processQueue, adjacentCoords)
					}
				}
			})
		}
	}

	return total
}

func isPaperRollAccessible(matrix WarehouseFloor, coords utils.Coordinates) bool {
	var adjacent int
	utils.EachSurroundingInMatrix(matrix, coords, func(adjacentObj string, _ utils.Coordinates, _ WarehouseFloor) {
		if adjacentObj == PaperRoll {
			adjacent++
		}
	})
	return adjacent <= MaxAdjacentToAccess
}

func removePaperRoll(room WarehouseFloor, coords utils.Coordinates) {
	utils.SetAtMatrixPosition(room, coords, EmptySpace)
}
