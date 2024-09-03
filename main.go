package main

import (
	year2023Day1 "advent_of_code/2023/01"
	year2023Day2 "advent_of_code/2023/02"
	year2023Day3 "advent_of_code/2023/03"
	utilities "advent_of_code/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	introduction()
	var year, day int
	userInput(&day, &year)
	solve(day, year)
}

func introduction() {
	fmt.Println("🎄", utilities.ColorText("green", "Advent of Code"), "🎅")
	printHr()
	fmt.Println(utilities.ColorText("cyan", "https://adventofcode.com/"))
	fmt.Print(utilities.ColorText("cyan", "Code By: Bryan Nikla"))
}

func userInput(day *int, year *int) {
	fmt.Println(utilities.ColorText("red", "\n\nEnter the year you want to run:"))
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	*year, _ = strconv.Atoi(input.Text())
	fmt.Println(utilities.ColorText("red", "\nEnter the day you want to run:"))
	input.Scan()
	*day, _ = strconv.Atoi(input.Text())
	fmt.Println()
}

func solve(day int, year int) {
	fmt.Println(year, "-", "Day", day)
	printHr()
	var solve1, solve2, test1, test2 = getSolution(year, day)
	printTestOutcome("Test 1", test1)
	printTestOutcome("Test 2", test2)
	printSolution("Solution 1", solve1)
	printSolution("Solution 2", solve2)
}

func getSolution(year int, day int) (any, any, bool, bool) {
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
			return nil, nil, false, false
		}
	default:
		return nil, nil, false, false
	}
}

func printTestOutcome(testName string, result bool) {
	if result {
		fmt.Println(testName+":", "✅")
	} else {
		fmt.Println(testName+":", "❌")
	}
}

func printSolution(name string, solution any) {
	fmt.Println(name+":", solution)
}

func printHr() {
	_, w := utilities.ConsoleSize()
	for range w {
		fmt.Print("-")
	}
	fmt.Print("\n")
}
