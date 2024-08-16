package main

import (
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

	Solve(day, year)
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
