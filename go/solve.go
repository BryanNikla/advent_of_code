package main

import (
	"advent_of_code/go/solutions/year2023"
	"fmt"
)

func Solve(day int, year int) {
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
