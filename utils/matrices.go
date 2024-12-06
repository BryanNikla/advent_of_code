package utils

import (
	"fmt"
	"slices"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func VisualizeMatrix[V any, M [][]V](matrix M, colored []Coordinates) {
	fmt.Print("\n")
	for yy, row := range matrix {
		for xx, element := range row {
			c := Coordinates{X: xx, Y: yy}
			if slices.Contains(colored, c) {
				fmt.Print(ColorText("red", element))
			} else {
				fmt.Print(element)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func CloneMatrix[V any](matrix [][]V) [][]V {
	newMatrix := make([][]V, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]V, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}

func SetAtMatrixPosition[V any, M [][]V](matrix M, coordinates Coordinates, value V) {
	matrix[coordinates.Y][coordinates.X] = value
}

func CallAtCords[V any, M [][]V](matrix M, cords Coordinates, fn func(V, Coordinates, M)) {
	if cords.X < 0 || cords.Y < 0 {
		return
	}
	if len(matrix) <= cords.Y {
		return
	}
	if len(matrix[cords.Y]) <= cords.X {
		return
	}

	fn(matrix[cords.Y][cords.X], cords, matrix)
}

func EachMatrix[V any, M [][]V](matrix M, fn func(V, Coordinates, M)) {
	for y, row := range matrix {
		for x := range row {
			cords := Coordinates{X: x, Y: y}
			fn(GetValueAtCords(matrix, cords), cords, matrix)
		}
	}
}

// EachSurroundingInMatrix - Calls function fn for every coordinate surrounding a set of cords in a matrix
func EachSurroundingInMatrix[Val any, M [][]Val](matrix M, cords Coordinates, fn func(Val, Coordinates, M)) {
	CallAtCords(matrix, Coordinates{X: cords.X, Y: cords.Y - 1}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X, Y: cords.Y + 1}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X - 1, Y: cords.Y}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X + 1, Y: cords.Y}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X - 1, Y: cords.Y - 1}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X + 1, Y: cords.Y + 1}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X - 1, Y: cords.Y + 1}, fn)
	CallAtCords(matrix, Coordinates{X: cords.X + 1, Y: cords.Y - 1}, fn)
}

func IsLastColOfMatrix[V any, M [][]V](matrix M, cords Coordinates) bool {
	return cords.Y == len(matrix[cords.X])-1
}

// GetValueAtCords - Return value present at matrix coordinates.
// Handles invalid coordinates gracefully by returning the default value for the expected type of value
func GetValueAtCords[V any, M [][]V](matrix M, cords Coordinates) V {
	defer func() V {
		recover()
		var defaultValue V
		return defaultValue
	}()
	return matrix[cords.Y][cords.X]
}

// LinesToCharacterMatrix - Return a matrix of individual string characters from a slice of strings (lines)
func LinesToCharacterMatrix(lines []string) [][]string {
	var matrix [][]string
	//for _, line := range lines {
	//	matrix = append(matrix, strings.Split(line, ""))
	//}
	// Do this in reverse so that if we ever 'visualize' this it makes sense // TODO: Make more sense of this
	for i := len(lines) - 1; i >= 0; i-- {
		matrix = append(matrix, strings.Split(lines[i], ""))
	}

	return matrix
}
