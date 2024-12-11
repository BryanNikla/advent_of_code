package year2024day11

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func Solve() utils.Solution {
	input1, input2 := utils.GetInputs(2024, 11)
	return utils.Solution{
		Day:   11,
		Test1: part1(input1) == 55312,
		Test2: part2(input2) == -1,
	}
}

func splitStringHalf(s string) (string, string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func blink(stones []string) []string {
	newStones := make([]string, 0, len(stones))
	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if len(stone)%2 == 0 {
			x, y := splitStringHalf(stone)
			newStones = append(newStones,
				strconv.Itoa(utils.StringToInteger(x)),
				strconv.Itoa(utils.StringToInteger(y)),
			)
		} else {
			newStones = append(newStones,
				strconv.Itoa(utils.StringToInteger(stone)*2024),
			)
		}
	}
	return newStones
}

func part1(input string) int {
	stones := strings.Split(input, " ")
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	fmt.Println("STONES:", len(stones))
	return len(stones)
}

func part2(input string) int {

	return 0
}
