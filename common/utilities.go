package utilities

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

func ColorText(color string, text string) string {
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
		return Red + text + Reset
	case "green":
		return Green + text + Reset
	case "yellow":
		return Yellow + text + Reset
	case "blue":
		return Blue + text + Reset
	case "magenta":
		return Magenta + text + Reset
	case "cyan":
		return Cyan + text + Reset
	case "gray":
		return Gray + text + Reset
	case "white":
		return White + text + Reset
	default:
		return text
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
	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return heigth, width
}

// Returns true if all elements in the slice satisfy the predicate, and false otherwise
// Basically mimics Javascript's Array.every() method
func Every[T any](slice []T, predicate func(T) bool) bool {
	for _, element := range slice {
		if !predicate(element) {
			return false
		}
	}
	return true
}
