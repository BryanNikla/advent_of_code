package main

import (
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
	"advent_of_code/year2025"
)

var solutionRegistry = map[int]map[int]func() utils.Solution{
	2023: {
		1: year2023Day1.Day1,
		2: year2023Day2.Day2,
		3: year2023Day3.Day3,
	},
	2024: {
		1:  year2024day01.Solve,
		2:  year2024day02.Solve,
		3:  year2024day03.Solve,
		4:  year2024day04.Solve,
		5:  year2024day05.Solve,
		6:  year2024day06.Solve,
		7:  year2024day07.Solve,
		8:  year2024day08.Solve,
		9:  year2024day09.Solve,
		10: year2024day10.Solve,
		11: year2024day11.Solve,
	},
	2025: {
		1: year2025.SolutionDay1,
		2: year2025.SolutionDay2,
		3: year2025.SolutionDay3,
		4: year2025.SolutionDay4,
		5: year2025.SolutionDay5,
	},
}
