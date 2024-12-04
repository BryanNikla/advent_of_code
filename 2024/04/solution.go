package year2024day04

import (
	"advent_of_code/utils"
)

func Solve() utils.Solution {
	input, testInput1, testInput2 := utils.GetAllInputs(2024, 4)
	return utils.Solution{
		Day:   4,
		Part1: part1(input),
		Part2: part2(input),
		Test1: part1(testInput1) == 18,
		Test2: part2(testInput2) == 9,
	}
}

func part1(input string) int {
	var count int
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	utils.EachMatrix(matrix, func(char string, x int, y int, matrix [][]string) {
		if char == "X" {
			// Test forwards along X axis
			if utils.GetValueAtCords(matrix, x+1, y) == "M" && utils.GetValueAtCords(matrix, x+2, y) == "A" && utils.GetValueAtCords(matrix, x+3, y) == "S" {
				count++
			}
			// Test backwards along X axis
			if utils.GetValueAtCords(matrix, x-1, y) == "M" && utils.GetValueAtCords(matrix, x-2, y) == "A" && utils.GetValueAtCords(matrix, x-3, y) == "S" {
				count++
			}
			// test Up along Y axis
			if utils.GetValueAtCords(matrix, x, y+1) == "M" && utils.GetValueAtCords(matrix, x, y+2) == "A" && utils.GetValueAtCords(matrix, x, y+3) == "S" {
				count++
			}
			// test down along Y axis
			if utils.GetValueAtCords(matrix, x, y-1) == "M" && utils.GetValueAtCords(matrix, x, y-2) == "A" && utils.GetValueAtCords(matrix, x, y-3) == "S" {
				count++
			}
			// test Diagonal north-east
			if utils.GetValueAtCords(matrix, x+1, y+1) == "M" && utils.GetValueAtCords(matrix, x+2, y+2) == "A" && utils.GetValueAtCords(matrix, x+3, y+3) == "S" {
				count++
			}
			// test Diagonal south-east
			if utils.GetValueAtCords(matrix, x+1, y-1) == "M" && utils.GetValueAtCords(matrix, x+2, y-2) == "A" && utils.GetValueAtCords(matrix, x+3, y-3) == "S" {
				count++
			}
			// test Diagonal south-west
			if utils.GetValueAtCords(matrix, x-1, y-1) == "M" && utils.GetValueAtCords(matrix, x-2, y-2) == "A" && utils.GetValueAtCords(matrix, x-3, y-3) == "S" {
				count++
			}
			// test Diagonal north-west
			if utils.GetValueAtCords(matrix, x-1, y+1) == "M" && utils.GetValueAtCords(matrix, x-2, y+2) == "A" && utils.GetValueAtCords(matrix, x-3, y+3) == "S" {
				count++
			}
		}
	})

	return count
}

func part2(input string) int {
	var count int
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	utils.EachMatrix(matrix, func(char string, x int, y int, matrix [][]string) {
		if char == "A" {
			// M.M
			// .A.
			// S.S
			if utils.GetValueAtCords(matrix, x-1, y+1) == "M" && utils.GetValueAtCords(matrix, x+1, y+1) == "M" {
				if utils.GetValueAtCords(matrix, x-1, y-1) == "S" && utils.GetValueAtCords(matrix, x+1, y-1) == "S" {
					count++
				}
			}
			// M.S
			// .A.
			// M.S
			if utils.GetValueAtCords(matrix, x-1, y+1) == "M" && utils.GetValueAtCords(matrix, x+1, y+1) == "S" {
				if utils.GetValueAtCords(matrix, x-1, y-1) == "M" && utils.GetValueAtCords(matrix, x+1, y-1) == "S" {
					count++
				}
			}
			// S.S
			// .A.
			// M.M
			if utils.GetValueAtCords(matrix, x-1, y+1) == "S" && utils.GetValueAtCords(matrix, x+1, y+1) == "S" {
				if utils.GetValueAtCords(matrix, x-1, y-1) == "M" && utils.GetValueAtCords(matrix, x+1, y-1) == "M" {
					count++
				}
			}
			// S.M
			// .A.
			// S.M
			if utils.GetValueAtCords(matrix, x-1, y+1) == "S" && utils.GetValueAtCords(matrix, x+1, y+1) == "M" {
				if utils.GetValueAtCords(matrix, x-1, y-1) == "S" && utils.GetValueAtCords(matrix, x+1, y-1) == "M" {
					count++
				}
			}
		}
	})

	return count
}
