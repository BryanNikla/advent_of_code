package year2024

import (
	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2024, 6, func() utils.Solution {
		input1, input2 := utils.GetInput(2024, 6)
		return utils.Solution{
			Day:   6,
			Test1: day6part1(input1) == 41,
			Test2: day6part2(input2) == 6,
		}
	})
}

func day6part1(input string) int {
	laboratory := utils.LinesToCharacterMatrix(utils.GetLines(input))
	guard := Guard{getStartingPositionOfGuard(laboratory), "N"}

	// Create map of every coordinate guard has been
	visited := make(map[utils.Coordinates]int)

	// Add initial guard position to visited map
	visited[guard.coordinates] = visited[guard.coordinates] + 1

	utils.SetAtMatrixPosition(laboratory, guard.coordinates, FLOOR)

	if err := utils.Forever(20000, func(exit func()) {
		coordinates := getNextPosition(&guard)
		switch utils.GetValueAtCords(laboratory, coordinates) {
		case EXIT:
			exit()
		case OBSTRUCTION:
			rotateGuard(&guard)
		case FLOOR:
			visited[coordinates] = visited[coordinates] + 1
			guard.coordinates = coordinates
		}
	}); err != nil {
		panic(err)
	}

	return len(visited)
}

func day6part2(input string) int {
	laboratory := utils.LinesToCharacterMatrix(utils.GetLines(input))
	guard := Guard{getStartingPositionOfGuard(laboratory), "N"}

	// hold positions to test, with value being true/false if it's a successful obstruction option
	positions := make(map[utils.Coordinates]bool)
	utils.EachMatrix(laboratory, func(content string, cords utils.Coordinates, _ [][]string) {
		if content == FLOOR {
			positions[cords] = true
		}
	})

	// Replace guard's starting position with traversable floor
	utils.SetAtMatrixPosition(laboratory, guard.coordinates, FLOOR)

	var count int
	for coordinates, _ := range positions {
		m := utils.CloneMatrix(laboratory)
		g := Guard{coordinates: utils.Coordinates{X: guard.coordinates.X, Y: guard.coordinates.Y}, facing: guard.facing}
		utils.SetAtMatrixPosition(m, coordinates, OBSTRUCTION)
		if looping := simulateGuard(&g, m); looping != nil {
			count++
		}
	}
	return count
}

type Guard struct {
	coordinates utils.Coordinates
	facing      string
}

var EXIT = ""
var FLOOR = "."
var VISITED = "@"
var OBSTRUCTION = "#"

// simulateGuard - Simulate the guard walking the laboratory.  Returns error if guard stuck in a loop
func simulateGuard(guard *Guard, lab [][]string) error {
	return utils.Forever(20000, func(exit func()) {
		coordinates := getNextPosition(guard)
		switch utils.GetValueAtCords(lab, coordinates) {
		case EXIT:
			exit()
		case OBSTRUCTION:
			rotateGuard(guard)
		case FLOOR:
			guard.coordinates = coordinates
		}
	})
}

// rotateGuard - Rotates guard clockwise from its current facing direction
func rotateGuard(guard *Guard) {
	switch guard.facing {
	case "N":
		guard.facing = "E"
	case "E":
		guard.facing = "S"
	case "S":
		guard.facing = "W"
	case "W":
		guard.facing = "N"
	}
}

func getNextPosition(guard *Guard) utils.Coordinates {
	switch guard.facing {
	case "N":
		return guard.coordinates.MoveNorth(1)
	case "E":
		return guard.coordinates.MoveEast(1)
	case "S":
		return guard.coordinates.MoveSouth(1)
	case "W":
		return guard.coordinates.MoveWest(1)
	default:
		panic("invalid guard position")
	}
}

func getStartingPositionOfGuard(laboratory [][]string) utils.Coordinates {
	var coordinates = utils.Coordinates{}
	utils.EachMatrix(laboratory, func(space string, cords utils.Coordinates, _ [][]string) {
		if space == "^" {
			coordinates = cords
		}
	})
	return coordinates
}
