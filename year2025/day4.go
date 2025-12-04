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
	PaperRoll           = '@'
	EmptySpace          = '.'
	MaxAdjacentToAccess = 3
)

func day4part1(input string) int {
	room := WarehouseFloor(utils.LinesToByteMatrix(utils.GetLines(input)))
	var totalAccessible int
	utils.EachMatrix(room, func(obj byte, coords utils.Coordinates, _ WarehouseFloor) {
		if obj == PaperRoll {
			if room.isPaperRollAccessible(coords) {
				totalAccessible++
			}
		}
	})
	return totalAccessible
}

func day4part2(input string) int {
	var totalRemoved int
	room := WarehouseFloor(utils.LinesToByteMatrix(utils.GetLines(input)))

	var processQueue []utils.Coordinates

	utils.EachMatrix(room, func(obj byte, coords utils.Coordinates, _ WarehouseFloor) {
		if obj == PaperRoll {
			if room.isPaperRollAccessible(coords) {
				processQueue = append(processQueue, coords)
			}
		}
	})

	for len(processQueue) > 0 {
		curr := processQueue[0]
		processQueue = processQueue[1:]

		if room.At(curr) == PaperRoll {
			room.removePaperRoll(curr)
			totalRemoved++
			utils.EachSurroundingInMatrix(room, curr, func(adjacentObj byte, adjacentCoords utils.Coordinates, _ WarehouseFloor) {
				if adjacentObj == PaperRoll {
					if room.isPaperRollAccessible(adjacentCoords) {
						processQueue = append(processQueue, adjacentCoords)
					}
				}
			})
		}
	}

	return totalRemoved
}

type WarehouseFloor [][]byte

func (f WarehouseFloor) isPaperRollAccessible(coords utils.Coordinates) bool {
	var adjacent int
	utils.EachSurroundingInMatrix(f, coords, func(adjacentObj byte, _ utils.Coordinates, _ WarehouseFloor) {
		if adjacentObj == PaperRoll {
			adjacent++
		}
	})
	return adjacent <= MaxAdjacentToAccess
}

func (f WarehouseFloor) removePaperRoll(coords utils.Coordinates) {
	utils.SetAtMatrixPosition(f, coords, EmptySpace)
}

func (f WarehouseFloor) At(coords utils.Coordinates) byte {
	return utils.GetValueAtCords(f, coords)
}
