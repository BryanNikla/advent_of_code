package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func GetFileContent(path string) string {
	file, _ := os.Open(path)
	b, _ := io.ReadAll(file)
	return string(b)
}

func GetInputs(year int, day int) (string, string) {
	return getInputContent(year, day, 1), getInputContent(year, day, 2)
}

func getInputContent(year int, day int, part int) string {
	var path = fmt.Sprintf("%d/%02d/part%d.txt", year, day, part)
	return GetFileContent(path)
}

func GetLines(input string) []string {
	return strings.Split(input, "\n")
}

// StringToInteger - Simpler string to integer that handles error (really just used to clean up solution logic)
func StringToInteger(input string) int {
	// Handle empty strings gracefully by returning 0
	if input == "" {
		return 0
	}
	integer, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return integer
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

// func ConsoleSize() (int, int) {
// 	cmd := exec.Command("stty", "size")
// 	cmd.Stdin = os.Stdin
// 	out, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	s := string(out)
// 	s = strings.TrimSpace(s)
// 	sArr := strings.Split(s, " ")
// 	height, err := strconv.Atoi(sArr[0])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	width, err := strconv.Atoi(sArr[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return height, width
// }

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

func GetRange(start int, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func ForRange(start int, end int, fn func(i int)) {
	for i := start; i <= end; i++ {
		fn(i)
	}
}

func InRange(value int, start int, end int) bool {
	return value >= start && value <= end
}
