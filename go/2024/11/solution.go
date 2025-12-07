package year2024day11

import (
	"strconv"
	"strings"

	utils "advent_of_code/utils"
)

func Solve() utils.Solution {
	input1, input2 := utils.GetInputs(2024, 11)
	return utils.Solution{
		Day:   11,
		Test1: part1(input1) == 55312,
		Test2: part2(input2) == 65601038650482,
	}
}

// part1 - Part 1 of this puzzle is to get the total number of stones after 25 "blinks"
func part1(input string) int {
	stones := createStoneMap(input)
	for range 25 {
		stones = blink(stones)
	}
	return countStones(stones)
}

// part2 - Same as part 1, but requiring 75 "blinks"
func part2(input string) int {
	stones := createStoneMap(input)
	for range 75 {
		stones = blink(stones)
	}
	return countStones(stones)
}

// blink - Simulate a single "blink" for a map of stones
// Returns a new map of stones with updated counts for each type of stone
func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
		} else if len(strconv.Itoa(stone))%2 == 0 {
			left, right := splitStone(stone)
			newStones[left] += count
			newStones[right] += count
		} else {
			newStones[stone*2024] += count
		}
	}
	return newStones
}

// splitStone - Given a stone's number, split it into two equally sized parts when interpretted as a string
// Returns two new numbers recast as an integer (removing any leading zeros potentially)
func splitStone(stone int) (int, int) {
	str := strconv.Itoa(stone)
	half := len(str) / 2
	l, r := str[:half], str[half:]
	return utils.StringToInteger(l), utils.StringToInteger(r)
}

// createStoneMap - Create a map of stone's with numbers & the total number of similar stones that exist
func createStoneMap(input string) map[int]int {
	stones := make(map[int]int)
	for _, s := range strings.Split(input, " ") {
		stones[utils.StringToInteger(s)] += 1
	}
	return stones
}

// countStones - Sum up the total number of stones across all stones in a stone map
func countStones(stones map[int]int) int {
	count := 0
	for _, v := range stones {
		count += v
	}
	return count
}
