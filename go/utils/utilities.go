package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

func GetFileContent(path string) string {
	file, _ := os.Open(path)
	b, _ := io.ReadAll(file)
	return string(b)
}

func GetInput(year int, day int) (string, string) {
	var p1 = fmt.Sprintf("inputs/y%d-d%d-p1.txt", year, day) // Part 1
	var p2 = fmt.Sprintf("inputs/y%d-d%d-p2.txt", year, day) // Part 2
	return GetFileContent(p1), GetFileContent(p2)
}

func GetLines(input string) []string {
	return strings.Split(input, "\n")
}

type ColorCycle struct {
	Color string
}

func NewColorCycle() *ColorCycle {
	return &ColorCycle{Color: "green"}
}

func (c *ColorCycle) NextColor() string {
	switch c.Color {
	case "green":
		c.Color = "red"
	case "red":
		c.Color = "yellow"
	case "yellow":
		c.Color = "blue"
	case "blue":
		c.Color = "magenta"
	case "magenta":
		c.Color = "cyan"
	case "cyan":
		c.Color = "white"
	case "white":
		fallthrough
	default:
		c.Color = "green"
	}
	return c.Color
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
	case "white":
		return White + asString + Reset
	default:
		return asString
	}
}

func ConsoleSize() (int, int) {
	// fd 1 is usually Stdout
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	// Your original function returned (height, width)
	return height, width
}

// Forever - Loop forever until a max iteration count is reached
func Forever(max int, fn func(f func())) error {
	var quit bool
	var iterations int
	for quit != true {
		iterations++
		if iterations > max {
			return errors.New("maximum iterations")
		}
		fn(func() {
			quit = true
		})
	}
	return nil
}

// AllPossibleCombinations generates all possible sequences of the specified length
// using elements from the provided slice. It effectively computes the Cartesian
// product of the items with themselves.
//
// Example: length=2, items=[1, 2]
// Returns: [[1, 1], [1, 2], [2, 1], [2, 2]]
//
// Note: Items can be reused, and order matters (e.g., [A, B] is distinct from [B, A]).
// Warning: The result size grows exponentially (len(items) ^ length).
func AllPossibleCombinations[V any](length int, items []V) [][]V {
	// Initialize a slice to store all combinations
	var allCombinations [][]V

	// Helper function to generate combinations recursively
	var recursiveFn func(curr []V)
	recursiveFn = func(curr []V) {
		// If we have the desired number of combinations, add to allCombinations
		if len(curr) == length {
			allCombinations = append(allCombinations, append([]V{}, curr...))
			return
		}

		// Recursively explore all possible operators for the next position
		for _, item := range items {
			curr = append(curr, item)
			recursiveFn(curr)
			curr = curr[:len(curr)-1] // Backtrack to the previous state
		}
	}

	// Start the backtracking process with an empty initial combination
	recursiveFn([]V{})

	return allCombinations
}

func GetRange[N Number](start N, end N) []N {
	var result []N
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func ForRange[N Number](start N, end N, fn func(i N)) {
	for i := start; i <= end; i++ {
		fn(i)
	}
}

func InRange[N Number](value N, start N, end N) bool {
	return value >= start && value <= end
}
