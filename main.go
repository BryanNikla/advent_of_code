package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"

	year2023Day1 "advent_of_code/2023/01"
	year2023Day2 "advent_of_code/2023/02"
	year2023Day3 "advent_of_code/2023/03"
	"advent_of_code/2024/go2024"
	"advent_of_code/utils"
)

func main() {
	introduction()
	var year, day int
	userInput(&day, &year)
	solve(day, year)
}

func introduction() {
	printHolidayHr()
	fmt.Println("🎄", utils.ColorText("green", "Advent of Code"), "🎅")
	printHolidayHr()
	fmt.Println(utils.ColorText("cyan", "https://adventofcode.com/"))
	fmt.Print(utils.ColorText("cyan", "Code By: Bryan Nikla"))
}

func userInput(day *int, year *int) {
	fmt.Println(utils.ColorText("red", "\n\nEnter the year you want to run:"))
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	*year, _ = strconv.Atoi(input.Text())
	fmt.Println(utils.ColorText("red", "\nEnter the day you want to run:"))
	input.Scan()
	*day, _ = strconv.Atoi(input.Text())
}

func solve(day int, year int) {
	var solution = getSolution(year, day)

	printFancyBox("Part 1", []string{
		formatTestOutcome("Test", solution.Test1),
		formatAnswer("Solution", solution.Part1),
	}, "blue")

	printFancyBox("Part 2", []string{
		formatTestOutcome("Test", solution.Test2),
		formatAnswer("Solution", solution.Part2),
	}, "magenta")
}

func getSolution(year int, day int) utils.Solution {
	switch year {
	case 2023:
		switch day {
		case 1:
			return year2023Day1.Day1()
		case 2:
			return year2023Day2.Day2()
		case 3:
			return year2023Day3.Day3()
		default:
			return utils.Solution{}
		}
	case 2024:
		return go2024.Execute(day)
	default:
		return utils.Solution{}
	}
}

func formatTestOutcome(testName string, result bool) string {
	if result {
		return formatAnswer(testName, "✅")
	} else {
		return formatAnswer(testName, "❌")
	}
}

func formatAnswer(name string, solution any) string {
	return fmt.Sprintf("%s: %v", name, solution)
}

func printHolidayHr() {
	_, width := utils.ConsoleSize()
	dashesToPrint := width - 5
	if dashesToPrint > 80 {
		dashesToPrint = 80
	}
	for i := 0; i < dashesToPrint; i++ {
		if i%2 == 0 {
			fmt.Print(utils.ColorText("red", "-"))
		} else {
			fmt.Print(utils.ColorText("green", "-"))
		}
	}
	fmt.Print("\n")
}

func printFancyBox(label string, contents []string, color string) {
	var width = 30 // minimum width of box
	for _, line := range contents {
		var length = utf8.RuneCountInString(line) + 5
		if length > width {
			width = length
		}
	}

	// Print the top label
	var labelDashCount = (width - utf8.RuneCountInString(label) - 2 - 2) / 2
	var labelDashes string
	for i := 0; i < labelDashCount; i++ {
		labelDashes += "-"
	}
	fmt.Println(utils.ColorText(color, "\n|"+labelDashes+" "+label+" "+labelDashes+"|"))

	// Print all contents
	for _, line := range contents {
		var whitespaceCount = width - utf8.RuneCountInString(line) - 2 - 3

		// TODO: Emojis do weird things for character counts/lengths when  printing. This is a dumb workaround for now.. solve one day maybe.
		var emojiDiff = len(line) - utf8.RuneCountInString(line)
		if emojiDiff > 0 {
			whitespaceCount = whitespaceCount - emojiDiff + 1
		}

		var whitespace string
		for i := 0; i < whitespaceCount; i++ {
			whitespace += " "
		}
		fmt.Println(utils.ColorText(color, "|"), line, whitespace, utils.ColorText(color, "|"))
	}

	// Print bottom of box
	var bottomDashCount = width - 2
	var bottomDashes string
	for i := 0; i < bottomDashCount; i++ {
		bottomDashes += "-"
	}
	fmt.Println(utils.ColorText(color, "|"+bottomDashes+"|"))
}
