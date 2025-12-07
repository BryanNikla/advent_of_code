package year2025

import (
	utils "advent_of_code/utils"
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
	floor.Each(func(obj byte, pos utils.Coordinates) {
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
	floor.Each(func(item byte, pos utils.Coordinates) {
		if item == PaperRoll {
			if floor.IsAccessible(pos) {
				queue = append(queue, pos)
			}
		}
	})
	for len(queue) > 0 {
		currPos := queue[0]
		queue = queue[1:]
		if floor.At(currPos) == PaperRoll {
			floor.Remove(currPos)
			totalRemoved++
			floor.EachNeighbor(currPos, func(item byte, nPos utils.Coordinates) {
				if item == PaperRoll {
					if floor.IsAccessible(nPos) {
						queue = append(queue, nPos)
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

func (floor WarehouseFloor) Each(fn func(obj byte, pos utils.Coordinates)) {
	utils.EachMatrix(floor, func(obj byte, pos utils.Coordinates, _ WarehouseFloor) {
		fn(obj, pos)
	})
}

func (floor WarehouseFloor) EachNeighbor(pos utils.Coordinates, fn func(obj byte, nPos utils.Coordinates)) {
	utils.EachSurroundingInMatrix(floor, pos, func(obj byte, nPos utils.Coordinates, _ WarehouseFloor) {
		fn(obj, nPos)
	})
}

func (floor WarehouseFloor) IsAccessible(pos utils.Coordinates) bool {
	var adjacent int
	floor.EachNeighbor(pos, func(adjacentItem byte, _ utils.Coordinates) {
		if adjacentItem == PaperRoll {
			adjacent++
		}
	})
	return adjacent <= MaxAdjacentToAccess
}

func (floor WarehouseFloor) Remove(pos utils.Coordinates) {
	utils.SetAtMatrixPosition(floor, pos, EmptySpace)
}

func (floor WarehouseFloor) At(pos utils.Coordinates) byte {
	return utils.GetValueAtCords(floor, pos)
}
