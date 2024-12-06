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
	year2024day01 "advent_of_code/2024/01"
	year2024day02 "advent_of_code/2024/02"
	year2024day03 "advent_of_code/2024/03"
	year2024day04 "advent_of_code/2024/04"
	year2024day05 "advent_of_code/2024/05"
	year2024day06 "advent_of_code/2024/06"
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
	for _, solution := range getSolutions(year, day) {
		printFancyBox(fmt.Sprintf("Day %d | Part 1", solution.Day), []string{
			formatTestOutcome("Test", solution.Test1),
			formatAnswer("Solution", solution.Part1),
		}, "blue")

		printFancyBox(fmt.Sprintf("Day %d | Part 2", solution.Day), []string{
			formatTestOutcome("Test", solution.Test2),
			formatAnswer("Solution", solution.Part2),
		}, "magenta")
	}
}

func getSolutions(year int, day int) []utils.Solution {
	switch year {
	case 2023:
		switch day {
		case 1:
			return []utils.Solution{year2023Day1.Day1()}
		case 2:
			return []utils.Solution{year2023Day2.Day2()}
		case 3:
			return []utils.Solution{year2023Day3.Day3()}
		default:
			return []utils.Solution{
				year2023Day1.Day1(),
				year2023Day2.Day2(),
				year2023Day3.Day3(),
			}
		}
	case 2024:
		switch day {
		case 1:
			return []utils.Solution{year2024day01.Solve()}
		case 2:
			return []utils.Solution{year2024day02.Solve()}
		case 3:
			return []utils.Solution{year2024day03.Solve()}
		case 4:
			return []utils.Solution{year2024day04.Solve()}
		case 5:
			return []utils.Solution{year2024day05.Solve()}
		case 6:
			return []utils.Solution{year2024day06.Solve()}
		default:
			return []utils.Solution{
				year2024day01.Solve(),
				year2024day02.Solve(),
				year2024day03.Solve(),
				year2024day04.Solve(),
				year2024day05.Solve(),
				//year2024day06.Solve(),
				{6, "---", "solution too slow for 'all'", true, true},
			}
		}
	default:
		return []utils.Solution{}
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
	var width = 50 // minimum width of box
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
