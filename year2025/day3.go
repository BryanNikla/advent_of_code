package year2025

import (
	"fmt"

	"advent_of_code/utils"
)

func SolutionDay3() utils.Solution {
	input := getInput(3)
	return utils.Solution{
		Day:   3,
		Test1: day3part1(input) == 357,
		Test2: day3part2(input) == 3121910778619,
	}
}

func day3part1(input string) int {
	totalOutput := 0
	batteryBanks := utils.GetLines(input)
	for _, bank := range batteryBanks {
		batteries := getBatteriesFromBank(bank)
		first, firstIndex := findLargestBattery(batteries[:len(batteries)-1])
		second, _ := findLargestBattery(batteries[firstIndex+1:])
		combined := fmt.Sprintf("%d%d", first, second)
		totalOutput += utils.StringToInteger(combined)
	}
	return totalOutput
}

func day3part2(input string) int {
	// Part 2 is something...... TODO
	return 0
}

func getBatteriesFromBank(bank string) []int {
	var batteries []int
	for _, char := range bank {
		battery := utils.StringToInteger(string(char))
		batteries = append(batteries, battery)
	}
	return batteries
}

func findLargestBattery(batteries []int) (int, int) {
	maxVal := batteries[0]
	maxIndex := 0
	for i, v := range batteries {
		if v > maxVal {
			maxIndex = i
			maxVal = v
		}
	}
	return maxVal, maxIndex
}
