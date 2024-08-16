package main

import (
	"advent_of_code/go/year2023"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	introduction()
	var year int
	var day int
	awaitInput(&day, &year)

	fmt.Println("Year:", year)
	fmt.Println("Day:", day)

	solve(day, year)
}

func introduction() {
	fmt.Println("🎄  Advent of Code  🎅")
	fmt.Println("----------------------")
	fmt.Println("https://adventofcode.com/")
	fmt.Print("Code By: Bryan Nikla")
}

func awaitInput(day *int, year *int) {
	fmt.Println("\nEnter the year you want to run:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	*year, _ = strconv.Atoi(input.Text())
	fmt.Println("\nEnter the day you want to run:")
	input.Scan()
	*day, _ = strconv.Atoi(input.Text())
}

func solve(day int, year int) {
	var part1, part2 = getSolution(year, day)
	fmt.Println(part1())
	fmt.Println(part2())
}

func getSolution(year int, day int) (func() string, func() string) {
	switch year {
	case 2023:
		switch day {
		case 1:
			return year2023.Day1()
		default:
			return func() string { return "" }, func() string { return "" }
		}
	default:
		return func() string { return "" }, func() string { return "" }
	}
}
