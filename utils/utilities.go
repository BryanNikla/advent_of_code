package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetFileContent(path string) string {
	file, _ := os.Open(path)
	b, _ := io.ReadAll(file)
	return string(b)
}

func GetAllInputs(year int, day int) (string, string, string) {
	return GetInputContent(year, day), GetTestContent(year, day, 1), GetTestContent(year, day, 2)
}

func GetInputContent(year int, day int) string {
	var path = fmt.Sprintf("%d/%02d/input.txt", year, day)
	return GetFileContent(path)
}

func GetTestContent(year int, day int, testNumber int) string {
	var path = fmt.Sprintf("%d/%02d/test%d.txt", year, day, testNumber)
	return GetFileContent(path)
}

func GetLines(input string) []string {
	return strings.Split(input, "\n")
}

func MiddleItem[V any](slice []V) V {
	middleIndex := len(slice) / 2
	return slice[middleIndex]
}

// StringToInteger - Simpler string to integer that handles error (really just used to clean up solution logic)
func StringToInteger(input string) int {
	integer, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return integer
}

// Reduce - Reducer for Slices. Initial value is always default for value type
func Reduce[T any, V any](slice []T, reducer func(accumulated V, currentValue T, currentIndex int) V) V {
	var value V
	for i := 0; i < len(slice); i++ {
		value = reducer(value, slice[i], i)
	}
	return value
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// AbsoluteValue - Helper function to calculate the absolute value
func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ColorText[V any](color string, text V) string {
	asString := fmt.Sprintf("%v", text)
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Cyan = "\033[36m"
	var Gray = "\033[37m"
	var White = "\033[97m"
	switch color {
	case "red":
		return Red + asString + Reset
	case "green":
		return Green + asString + Reset
	case "yellow":
		return Yellow + asString + Reset
	case "blue":
		return Blue + asString + Reset
	case "magenta":
		return Magenta + asString + Reset
	case "cyan":
		return Cyan + asString + Reset
	case "gray":
		return Gray + asString + Reset
	case "white":
		return White + asString + Reset
	default:
		return asString
	}
}

func ConsoleSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")
	height, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return height, width
}

// Every - Returns true if all elements in the slice satisfy the predicate, and false otherwise
// Basically mimics Javascript's Array.every() method
func Every[T any](slice []T, predicate func(T) bool) bool {
	for _, element := range slice {
		if !predicate(element) {
			return false
		}
	}
	return true
}

func SumValuesInSlice[V int | int8 | int16 | int32 | int64 | float32 | float64, S []V](slice S) V {
	var sum V
	for _, value := range slice {
		sum += value
	}
	return sum
}

func VisualizeMatrix[V any, M [][]V](matrix M, x int, y int) {
	fmt.Print("\n")
	for yy, row := range matrix {
		for xx, element := range row {
			if yy == y && xx == x {
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

func SetAtMatrixPosition[V any, M [][]V](matrix M, x int, y int, value V) {
	matrix[y][x] = value
}

func CallAtCords[V any, M [][]V](matrix M, x int, y int, fn func(V, int, int, M)) {
	if x < 0 || y < 0 {
		return
	}
	if len(matrix) <= x {
		return
	}
	if len(matrix[x]) <= y {
		return
	}

	fn(matrix[x][y], x, y, matrix)
}

func EachMatrix[V any, M [][]V](matrix M, fn func(V, int, int, M)) {
	for y, row := range matrix {
		for x := range row {
			fn(GetValueAtCords(matrix, x, y), x, y, matrix)
		}
	}
}

// EachSurroundingInMatrix - Calls function fn for every coordinate surrounding a set of cords in a matrix
func EachSurroundingInMatrix[Val any, M [][]Val](matrix M, x int, y int, fn func(Val, int, int, M)) {
	CallAtCords(matrix, x, y-1, fn)
	CallAtCords(matrix, x, y+1, fn)
	CallAtCords(matrix, x-1, y, fn)
	CallAtCords(matrix, x+1, y, fn)
	CallAtCords(matrix, x-1, y-1, fn)
	CallAtCords(matrix, x+1, y+1, fn)
	CallAtCords(matrix, x-1, y+1, fn)
	CallAtCords(matrix, x+1, y-1, fn)
}

func IsLastColOfMatrix[V any, M [][]V](matrix M, x int, y int) bool {
	return y == len(matrix[x])-1
}

// GetValueAtCords - Return value present at matrix coordinates.
// Handles invalid coordinates gracefully by returning the default value for the expected type of value
func GetValueAtCords[V any, M [][]V](matrix M, x int, y int) V {
	defer func() V {
		recover()
		var defaultValue V
		return defaultValue
	}()
	return matrix[y][x]
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
