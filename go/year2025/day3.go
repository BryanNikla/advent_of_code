package year2025

import (
	utils "advent_of_code/utils"
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
	output := 0
	for _, bank := range utils.GetLines(input) {
		batteries := parseBatteries(bank)
		first, firstIdx := findLargestBattery(batteries[:len(batteries)-1])
		second, _ := findLargestBattery(batteries[firstIdx+1:])
		output += (first * 10) + second
	}
	return output
}

func day3part2(input string) int {
	output := 0
	for _, bank := range utils.GetLines(input) {
		batteries := parseBatteries(bank)
		batteryCount := len(batteries)
		bankValue := 0
		searchStartIdx := 0
		for remaining := 11; remaining >= 0; remaining-- {
			battery, relativeIdx := findLargestBattery(batteries[searchStartIdx : batteryCount-remaining])
			searchStartIdx += relativeIdx + 1
			bankValue = (bankValue * 10) + battery
		}
		output += bankValue
	}
	return output
}

// parseBatteries converts a string of digits into a slice of integers
func parseBatteries(bank string) []int {
	ret := make([]int, len(bank))
	for i, char := range bank {
		ret[i] = utils.StringToInteger(string(char))
	}
	return ret
}

// findLargestBattery returns max battery value & its relative index in the slice
func findLargestBattery(batteries []int) (int, int) {
	maxBattery := batteries[0]
	maxIndex := 0
	for i, v := range batteries {
		if v > maxBattery {
			maxIndex = i
			maxBattery = v
		}
	}
	return maxBattery, maxIndex
}
