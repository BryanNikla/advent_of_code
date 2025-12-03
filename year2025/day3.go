package year2025

import (
	"fmt"
	"strconv"
	"strings"

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
		concatenated := fmt.Sprintf("%d%d", first, second)
		totalOutput += utils.StringToInteger(concatenated)
	}
	return totalOutput
}

func day3part2(input string) int {
	totalOutput := 0
	batteryBanks := utils.GetLines(input)
	for _, bank := range batteryBanks {
		allBatteries := getBatteriesFromBank(bank)
		totalLength := len(allBatteries)
		batteriesFound := make([]string, 0, 12)
		var battery, lastIdx, currIdx int
		for i := 11; i > -1; i-- {
			battery, currIdx = findLargestBattery(allBatteries[lastIdx : totalLength-i])
			lastIdx += currIdx + 1
			batteriesFound = append(batteriesFound, strconv.Itoa(battery))
		}
		concatenated := strings.Join(batteriesFound, "")
		totalOutput += utils.StringToInteger(concatenated)
	}
	return totalOutput
}

func getBatteriesFromBank(bank string) []int {
	var batteries []int
	for _, char := range bank {
		battery := utils.StringToInteger(string(char))
		batteries = append(batteries, battery)
	}
	return batteries
}

// Find the largest battery for a given slice of batteries, returning its value & index
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
