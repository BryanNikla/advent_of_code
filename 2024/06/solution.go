package year2024day06

import (
	"errors"

	"advent_of_code/utils"
)

func Solve() utils.Solution {
	input, testInput1, testInput2 := utils.GetAllInputs(2024, 6)
	return utils.Solution{
		Day:   6,
		Part1: part1(input),
		Part2: part2(input),
		Test1: part1(testInput1) == 41,
		Test2: part2(testInput2) == 6,
	}
}

func part1(input string) int {
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	guard := Guard{getStartingPositionOfGuard(matrix), "N"}

	visited := make(map[utils.Coordinates]int)
	addToVisitedMap := func(coordinates utils.Coordinates) {
		visited[coordinates] = visited[coordinates] + 1
	}

	addToVisitedMap(guard.coordinates)
	utils.SetAtMatrixPosition(matrix, guard.coordinates, FLOOR)

	var exited bool
	var iterations int
	for exited != true {
		coordinates := getNextPosition(&guard)
		switch utils.GetValueAtCords(matrix, coordinates) {
		case EXIT:
			exited = true
		case OBSTRUCTION:
			rotateGuard(&guard)
		case FLOOR:
			addToVisitedMap(coordinates)
			guard.coordinates = coordinates
		}
		if iterations++; iterations > 20000 {
			panic("something is very wrong: stuck")
		}
	}

	return len(visited)
}

func part2(input string) int {
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	guard := Guard{getStartingPositionOfGuard(matrix), "N"}

	// hold positions to test, with value being true/false if it's a successful obstruction option
	positions := make(map[utils.Coordinates]bool)
	utils.EachMatrix(matrix, func(content string, cords utils.Coordinates, _ [][]string) {
		if content == FLOOR {
			positions[cords] = true
		}
	})

	utils.SetAtMatrixPosition(matrix, guard.coordinates, VISITED)

	var count int
	for coordinates, _ := range positions {
		m := utils.CloneMatrix(matrix)
		g := Guard{coordinates: utils.Coordinates{X: guard.coordinates.X, Y: guard.coordinates.Y}, facing: guard.facing}
		utils.SetAtMatrixPosition(m, coordinates, OBSTRUCTION)
		if looping := simulateGuard(&g, m); looping != nil {
			count++
		}
	}
	return count
}

var EXIT = ""
var FLOOR = "."
var VISITED = "@"
var OBSTRUCTION = "#"

// simulateGuard - Simulate the guard walking the matrix. Will return error if guard is looping
func simulateGuard(guard *Guard, matrix [][]string) error {
	var exited bool
	var iterations int
	for exited != true {
		iterations++
		// TODO: Brute forcing this solution by simulating.. this is NOT ideal (but it worked)
		// TODO: Maybe come back and revisit and solve a better way one day
		if iterations > 20000 {
			return errors.New("guard is looping")
		}
		utils.SetAtMatrixPosition(matrix, guard.coordinates, VISITED)
		coordinates := getNextPosition(guard)
		switch utils.GetValueAtCords(matrix, coordinates) {
		case EXIT:
			exited = true
		case OBSTRUCTION:
			rotateGuard(guard)
		case VISITED:
			fallthrough
		case FLOOR:
			utils.SetAtMatrixPosition(matrix, coordinates, VISITED)
			guard.coordinates = coordinates
		}
	}
	return nil
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
		return utils.Coordinates{X: guard.coordinates.X, Y: guard.coordinates.Y + 1}
	case "E":
		return utils.Coordinates{X: guard.coordinates.X + 1, Y: guard.coordinates.Y}
	case "S":
		return utils.Coordinates{X: guard.coordinates.X, Y: guard.coordinates.Y - 1}
	case "W":
		return utils.Coordinates{X: guard.coordinates.X - 1, Y: guard.coordinates.Y}
	default:
		panic("invalid guard position")
	}
}

func getStartingPositionOfGuard(matrix [][]string) utils.Coordinates {
	var coordinates = utils.Coordinates{}
	utils.EachMatrix(matrix, func(space string, cords utils.Coordinates, _ [][]string) {
		if space == "^" {
			coordinates = cords
		}
	})
	return coordinates
}

type Guard struct {
	coordinates utils.Coordinates
	facing      string
}
