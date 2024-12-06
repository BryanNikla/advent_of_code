package year2024day04

import (
	"advent_of_code/utils"
)

func Solve() utils.Solution {
	testInput1, testInput2 := utils.GetInputs(2024, 4)
	return utils.Solution{
		Day:   4,
		Test1: part1(testInput1) == 18,
		Test2: part2(testInput2) == 9,
	}
}

func part1(input string) int {
	var count int
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	utils.EachMatrix(matrix, func(char string, cords utils.Coordinates, matrix [][]string) {
		if char == "X" {
			// Test forwards along X axis
			if utils.GetValueAtCords(matrix, cords.MoveEast(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveEast(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveEast(3)) == "S" {
				count++
			}
			// Test backwards along X axis
			if utils.GetValueAtCords(matrix, cords.MoveWest(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveWest(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveWest(3)) == "S" {
				count++
			}
			// test Up along Y axis
			if utils.GetValueAtCords(matrix, cords.MoveNorth(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveNorth(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveNorth(3)) == "S" {
				count++
			}
			// test down along Y axis
			if utils.GetValueAtCords(matrix, cords.MoveSouth(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveSouth(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveSouth(3)) == "S" {
				count++
			}
			// test Diagonal north-east
			if utils.GetValueAtCords(matrix, cords.MoveNorthEast(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthEast(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthEast(3)) == "S" {
				count++
			}
			// test Diagonal south-east
			if utils.GetValueAtCords(matrix, cords.MoveSouthEast(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveSouthEast(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveSouthEast(3)) == "S" {
				count++
			}
			// test Diagonal south-west
			if utils.GetValueAtCords(matrix, cords.MoveSouthWest(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveSouthWest(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveSouthWest(3)) == "S" {
				count++
			}
			// test Diagonal north-west
			if utils.GetValueAtCords(matrix, cords.MoveNorthWest(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthWest(2)) == "A" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthWest(3)) == "S" {
				count++
			}
		}
	})

	return count
}

func part2(input string) int {
	var count int
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	utils.EachMatrix(matrix, func(char string, cords utils.Coordinates, matrix [][]string) {
		if char == "A" {
			// M.M
			// .A.
			// S.S
			if utils.GetValueAtCords(matrix, cords.MoveNorthWest(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthEast(1)) == "M" {
				if utils.GetValueAtCords(matrix, cords.MoveSouthWest(1)) == "S" &&
					utils.GetValueAtCords(matrix, cords.MoveSouthEast(1)) == "S" {
					count++
				}
			}
			// M.S
			// .A.
			// M.S
			if utils.GetValueAtCords(matrix, cords.MoveNorthWest(1)) == "M" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthEast(1)) == "S" {
				if utils.GetValueAtCords(matrix, cords.MoveSouthWest(1)) == "M" &&
					utils.GetValueAtCords(matrix, cords.MoveSouthEast(1)) == "S" {
					count++
				}
			}
			// S.S
			// .A.
			// M.M
			if utils.GetValueAtCords(matrix, cords.MoveNorthWest(1)) == "S" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthEast(1)) == "S" {
				if utils.GetValueAtCords(matrix, cords.MoveSouthWest(1)) == "M" &&
					utils.GetValueAtCords(matrix, cords.MoveSouthEast(1)) == "M" {
					count++
				}
			}
			// S.M
			// .A.
			// S.M
			if utils.GetValueAtCords(matrix, cords.MoveNorthWest(1)) == "S" &&
				utils.GetValueAtCords(matrix, cords.MoveNorthEast(1)) == "M" {
				if utils.GetValueAtCords(matrix, cords.MoveSouthWest(1)) == "S" &&
					utils.GetValueAtCords(matrix, cords.MoveSouthEast(1)) == "M" {
					count++
				}
			}
		}
	})

	return count
}
