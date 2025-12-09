package year2024

import (
	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2024, 8, func() utils.Solution {
		input1, input2 := utils.GetInput(2024, 8)
		return utils.Solution{
			Day:   8,
			Test1: day8part1(input1) == 14,
			Test2: day8part2(input2) == 34,
		}
	})
}

func day8part1(input string) int {
	frequencyMap := utils.LinesToCharacterMatrix(utils.GetLines(input))

	// Map to hold all distinct antinodes found
	antinodes := make(map[utils.Coordinates]bool)

	utils.EachMatrix(frequencyMap, func(char1 string, cords1 utils.Coordinates, _ [][]string) {
		if char1 != "." {
			utils.EachMatrix(frequencyMap, func(char2 string, cords2 utils.Coordinates, _ [][]string) {
				if cords1 != cords2 && char1 == char2 {
					possibleAntinode := findThirdCoordinate(&cords1, &cords2)
					if utils.GetValueAtCords(frequencyMap, possibleAntinode) != "" {
						antinodes[possibleAntinode] = true
					}
				}
			})
		}
	})
	return len(antinodes)
}

func day8part2(input string) int {
	frequencyMap := utils.LinesToCharacterMatrix(utils.GetLines(input))

	// Map to hold all distinct antinodes found
	antinodes := make(map[utils.Coordinates]bool)

	utils.EachMatrix(frequencyMap, func(char1 string, cords1 utils.Coordinates, _ [][]string) {
		if char1 != "." {
			utils.EachMatrix(frequencyMap, func(char2 string, cords2 utils.Coordinates, _ [][]string) {
				if cords1 != cords2 && char1 == char2 {
					// Since these two line up already, they are themselves already antinodes
					antinodes[cords1] = true
					antinodes[cords2] = true

					var checkRecursively func(c1 *utils.Coordinates, c2 *utils.Coordinates)
					checkRecursively = func(c1 *utils.Coordinates, c2 *utils.Coordinates) {
						possibleAntinode := findThirdCoordinate(c1, c2)
						if utils.GetValueAtCords(frequencyMap, possibleAntinode) != "" {
							antinodes[possibleAntinode] = true
							checkRecursively(c2, &possibleAntinode)
						}
					}
					checkRecursively(&cords1, &cords2)
				}
			})
		}
	})
	return len(antinodes)
}

func findThirdCoordinate(c1, c2 *utils.Coordinates) utils.Coordinates {
	// Calculate the third point by the distance between each x & y to coordinate2
	return utils.Coordinates{
		X: c2.X + (c2.X - c1.X),
		Y: c2.Y + (c2.Y - c1.Y),
	}
}
