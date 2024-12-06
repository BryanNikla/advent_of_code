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
