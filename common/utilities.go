package utilities

import (
	"fmt"
	"io"
	"os"
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
