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
	var totalAccessible int
	floor := NewWarehouseFloor(input)
	utils.EachMatrix(floor, func(obj byte, pos utils.Coordinates, _ WarehouseFloor) {
		if obj == PaperRoll {
			if floor.IsAccessible(pos) {
				totalAccessible++
			}
		}
	})
	return totalAccessible
}

func day4part2(input string) int {
	var totalRemoved int
	floor := NewWarehouseFloor(input)
	var queue []utils.Coordinates
	utils.EachMatrix(floor, func(item byte, pos utils.Coordinates, _ WarehouseFloor) {
		if item == PaperRoll {
			if floor.IsAccessible(pos) {
				queue = append(queue, pos)
			}
		}
	})
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if floor.At(curr) == PaperRoll {
			floor.Remove(curr)
			totalRemoved++
			utils.EachSurroundingInMatrix(floor, curr, func(item byte, pos utils.Coordinates, _ WarehouseFloor) {
				if item == PaperRoll {
					if floor.IsAccessible(pos) {
						queue = append(queue, pos)
					}
				}
			})
		}
	}
	return totalRemoved
}

type WarehouseFloor [][]byte

func NewWarehouseFloor(input string) WarehouseFloor {
	return utils.LinesToByteMatrix(utils.GetLines(input))
}

func (f WarehouseFloor) IsAccessible(pos utils.Coordinates) bool {
	var adjacent int
	utils.EachSurroundingInMatrix(f, pos, func(adjacentItem byte, _ utils.Coordinates, _ WarehouseFloor) {
		if adjacentItem == PaperRoll {
			adjacent++
		}
	})
	return adjacent <= MaxAdjacentToAccess
}

func (f WarehouseFloor) Remove(pos utils.Coordinates) {
	utils.SetAtMatrixPosition(f, pos, EmptySpace)
}

func (f WarehouseFloor) At(pos utils.Coordinates) byte {
	return utils.GetValueAtCords(f, pos)
}
