package year2024day06

import (
	"errors"
	"fmt"

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
	guard := Guard{facing: "N"}

	visited := make(map[Coordinates]int)

	addToVisitedMap := func(x int, y int) {
		visited[Coordinates{x, y}] = visited[Coordinates{x, y}] + 1
	}

	// Get starting position of guard
	startCoordinates := getStartingPositionOfGuard(matrix)
	guard.x = startCoordinates.x
	guard.y = startCoordinates.y
	addToVisitedMap(guard.x, guard.y)
	utils.SetAtMatrixPosition(matrix, guard.x, guard.y, VISITED)

	var exited bool
	for exited != true {
		var content, x, y = getNextPosition(matrix, &guard)
		switch content {
		case EXIT:
			exited = true
		case OBSTRUCTION:
			rotateGuard(&guard)
		case VISITED:
			fallthrough
		case FLOOR:
			addToVisitedMap(x, y)
			utils.SetAtMatrixPosition(matrix, x, y, VISITED)
			guard.x = x
			guard.y = y
		}
	}

	return len(visited)
}

func part2(input string) int {
	matrix := utils.LinesToCharacterMatrix(utils.GetLines(input))
	guard := Guard{facing: "N"}

	// hold positions to test, with value being true/false if it's a successful obstruction option
	positions := make(map[Coordinates]bool)

	// Get starting position of guard
	startCoordinates := getStartingPositionOfGuard(matrix)
	guard.x = startCoordinates.x
	guard.y = startCoordinates.y

	// Populate positions map
	utils.EachMatrix(matrix, func(content string, x int, y int, _ [][]string) {
		if content == FLOOR {
			positions[Coordinates{x: x, y: y}] = true
		}
	})

	utils.SetAtMatrixPosition(matrix, guard.x, guard.y, VISITED)

	var count int

	for coordinates, _ := range positions {
		m := utils.CloneMatrix(matrix)
		g := Guard{guard.x, guard.y, guard.facing}
		utils.SetAtMatrixPosition(m, coordinates.x, coordinates.y, OBSTRUCTION)

		if looping := simulateGuard(&g, m); looping != nil {
			count++
		}
	}

	fmt.Println("Number of Successful Obstructions:", count)
	return count
}

// simulateGuard - Simulate the guard walking the matrix. Will return error if guard is looping
func simulateGuard(guard *Guard, matrix [][]string) error {
	var exited bool
	var iterations int
	for exited != true {
		iterations++
		if iterations > 20000 {
			return errors.New("guard is looping")
		}
		utils.SetAtMatrixPosition(matrix, guard.x, guard.y, VISITED)
		var content, x, y = getNextPosition(matrix, guard)
		switch content {
		case EXIT:
			exited = true
		case OBSTRUCTION:
			rotateGuard(guard)
		case VISITED:
			fallthrough
		case FLOOR:
			utils.SetAtMatrixPosition(matrix, x, y, VISITED)
			guard.x = x
			guard.y = y
		}
	}
	return nil
}

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

func getNextPosition(matrix [][]string, guard *Guard) (string, int, int) {
	switch guard.facing {
	case "N":
		return utils.GetValueAtCords(matrix, guard.x, guard.y+1), guard.x, guard.y + 1
	case "E":
		return utils.GetValueAtCords(matrix, guard.x+1, guard.y), guard.x + 1, guard.y
	case "S":
		return utils.GetValueAtCords(matrix, guard.x, guard.y-1), guard.x, guard.y - 1
	case "W":
		return utils.GetValueAtCords(matrix, guard.x-1, guard.y), guard.x - 1, guard.y
	default:
		return "", 0, 0
	}
}

func getStartingPositionOfGuard(matrix [][]string) Coordinates {
	var coordinates = Coordinates{-1, -1}
	utils.EachMatrix(matrix, func(space string, x int, y int, _ [][]string) {
		if space == "^" {
			coordinates.x = x
			coordinates.y = y
		}
	})
	return coordinates
}

var OBSTRUCTION = "#"
var FLOOR = "."
var EXIT = ""
var VISITED = "@"

type Guard struct {
	x      int
	y      int
	facing string
}

type Coordinates struct {
	x int
	y int
}
