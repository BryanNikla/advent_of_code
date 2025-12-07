package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"advent_of_code/registry"
	"advent_of_code/utils"

	// Import year packages to register them
	_ "advent_of_code/year2023"
	_ "advent_of_code/year2024"
	_ "advent_of_code/year2025"
)

var DefaultYear = time.Now().Year()

func main() {
	year := flag.Int("year", -1, "The year of AoC to run")
	day := flag.Int("day", -1, "The specific day to run (0 for all)")
	flag.Parse()
	introduction()
	userInput(day, year)
	solve(*day, *year)
}

func introduction() {
	printHolidayHr()
	fmt.Println("🎄", utils.ColorText("green", "Advent of Code"), "🎅")
	printHolidayHr()
	fmt.Println(utils.ColorText("cyan", "https://adventofcode.com/"))
	fmt.Print(utils.ColorText("cyan", "Code By: Bryan Nikla"))
	fmt.Println()
}

func userInput(day *int, year *int) {
	reader := bufio.NewReader(os.Stdin)
	if *year == -1 {
		fmt.Println(utils.ColorText("red", fmt.Sprintf("\n\nEnter the year (default %d):", DefaultYear)))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			*year = DefaultYear
		} else {
			y, err := strconv.Atoi(text)
			if err != nil {
				printIssue(fmt.Sprintf("Invalid input '%s', using default %d", text, DefaultYear))
				*year = DefaultYear
			} else {
				*year = y
			}
		}
	}
	if *day == -1 {
		fmt.Println(utils.ColorText("red", "\nEnter the day (default All):"))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		var defaultDay = 0 // 0 means all days
		if text == "" {
			*day = defaultDay
		} else {
			d, err := strconv.Atoi(text)
			if err != nil {
				printIssue(fmt.Sprintf("Invalid input '%s', running all days", text))
				*day = defaultDay
			} else {
				*day = d
			}
		}
	}
}

func solve(day int, year int) {
	colorCycle := utils.NewColorCycle()
	for _, solution := range registry.Registry.GetSolutions(year, day) {
		printContents(colorCycle.NextColor(), []string{
			fmt.Sprintf("Day %d", solution.Day),
			fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, solution.Day),
			formatSolutionOutcome(solution, 1),
			formatSolutionOutcome(solution, 2),
		})
	}
}

func formatSolutionOutcome(solution utils.Solution, part int) string {
	success := (part == 1 && solution.Test1) || (part == 2 && solution.Test2)
	icons := map[bool]string{true: "✅", false: "❌"}
	return fmt.Sprintf("Part %d: %v", part, icons[success])
}

func printHolidayHr() {
	_, width := utils.ConsoleSize()
	dashes := width - 5
	if dashes > 80 {
		dashes = 80
	}
	var sb strings.Builder
	for i := 0; i < dashes; i++ {
		if i%2 == 0 {
			sb.WriteString(utils.ColorText("red", "-"))
		} else {
			sb.WriteString(utils.ColorText("green", "-"))
		}
	}
	fmt.Println(sb.String())
}

func printContents(color string, contents []string) {
	fmt.Println()
	for _, line := range contents {
		fmt.Println(utils.ColorText(color, "|"), utils.ColorText(color, line))
	}
}

func printIssue(message string) {
	fmt.Println(utils.ColorText("red", fmt.Sprintf("\n⚠️ %s", message)))
}
