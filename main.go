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
	year2024day07 "advent_of_code/2024/07"
	year2024day08 "advent_of_code/2024/08"
	year2024day09 "advent_of_code/2024/09"
	year2024day10 "advent_of_code/2024/10"
	year2024day11 "advent_of_code/2024/11"
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
	colorCycle := utils.ColorCycle{}
	for _, solution := range getSolutions(year, day) {
		printFancy(fmt.Sprintf("Day %d", solution.Day), []string{
			fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, solution.Day),
			formatTestOutcome("Part 1", solution.Test1),
			formatTestOutcome("Part 2", solution.Test2),
		}, colorCycle.NextColor())
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
		case 7:
			return []utils.Solution{year2024day07.Solve()}
		case 8:
			return []utils.Solution{year2024day08.Solve()}
		case 9:
			return []utils.Solution{year2024day09.Solve()}
		case 10:
			return []utils.Solution{year2024day10.Solve()}
		case 11:
			return []utils.Solution{year2024day11.Solve()}
		default:
			return []utils.Solution{
				year2024day01.Solve(),
				year2024day02.Solve(),
				year2024day03.Solve(),
				year2024day04.Solve(),
				year2024day05.Solve(),
				year2024day06.Solve(),
				year2024day07.Solve(),
				year2024day08.Solve(),
				year2024day09.Solve(),
				year2024day10.Solve(),
				year2024day11.Solve(),
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

func printFancy(label string, contents []string, color string) {
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
	fmt.Println(utils.ColorText(color, "\n|"), utils.ColorText(color, label))

	// Print all contents
	for _, line := range contents {
		fmt.Println(utils.ColorText(color, "|"), utils.ColorText(color, line))
	}
}
