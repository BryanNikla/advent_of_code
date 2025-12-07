package year2024day10

import (
	utils "advent_of_code/utils"
)

func Solve() utils.Solution {
	input1, input2 := utils.GetInputs(2024, 10)
	return utils.Solution{
		Day:   10,
		Test1: part1(input1) == 36,
		Test2: part2(input2) == 81,
	}
}

func part1(input string) int {
	topographicMap := getTopographicMap(input)

	trailheads := make(map[utils.Coordinates]map[utils.Coordinates]bool)

	var checkPos func(trailhead utils.Coordinates, cords utils.Coordinates)
	checkPos = func(trailhead utils.Coordinates, cords utils.Coordinates) {
		x := utils.GetValueAtCords(topographicMap, cords)
		utils.EachSurroundingCardinalInMatrix(topographicMap, cords, func(xx int, c utils.Coordinates, _ [][]int) {
			if xx == x+1 {
				if xx == 9 {
					trailheads[trailhead][c] = true
				} else {
					checkPos(trailhead, c)
				}
			}
		})
	}

	// Find all trailheads and check their trails
	utils.EachMatrix(topographicMap, func(c int, cords utils.Coordinates, m [][]int) {
		if c == 0 {
			trailheads[cords] = make(map[utils.Coordinates]bool)
			checkPos(cords, cords)
		}
	})

	var sum int
	for _, x := range trailheads {
		sum = sum + len(x)
	}
	return sum
}

func part2(input string) int {
	topographicMap := getTopographicMap(input)

	var sum int

	var checkPos func(trailhead utils.Coordinates, cords utils.Coordinates)
	checkPos = func(trailhead utils.Coordinates, cords utils.Coordinates) {
		x := utils.GetValueAtCords(topographicMap, cords)
		utils.EachSurroundingCardinalInMatrix(topographicMap, cords, func(xx int, c utils.Coordinates, _ [][]int) {
			if xx == x+1 {
				if xx == 9 {
					sum++
				} else {
					checkPos(trailhead, c)
				}
			}
		})
	}

	// Find all trailheads and check their trails
	utils.EachMatrix(topographicMap, func(c int, cords utils.Coordinates, m [][]int) {
		if c == 0 {
			checkPos(cords, cords)
		}
	})

	return sum
}

func getTopographicMap(input string) [][]int {
	stringMap := utils.LinesToCharacterMatrix(utils.GetLines(input))
	topographicMap := make([][]int, len(stringMap))
	for i := range stringMap {
		topographicMap[i] = make([]int, len(stringMap[i]))
	}
	utils.EachMatrix(stringMap, func(c string, cords utils.Coordinates, m [][]string) {
		utils.SetAtMatrixPosition(topographicMap, cords, utils.StringToInteger(c))
	})
	return topographicMap
}
