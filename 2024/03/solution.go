package year2024day03

import (
	"regexp"
	"strings"

	"advent_of_code/utils"
)

func Solve() utils.Solution {
	input, testInput1, testInput2 := utils.GetAllInputs(2024, 3)
	return utils.Solution{
		Day:   3,
		Part1: part1(input),
		Part2: part2(input),
		Test1: part1(testInput1) == 161,
		Test2: part2(testInput2) == 48,
	}
}

func part1(input string) int {
	r := regexp.MustCompile(`mul\(([0-9]*),([0-9]*)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	return utils.Reduce(matches, func(total int, match []string, _ int) int {
		return total + (utils.StringToInteger(match[1]) * utils.StringToInteger(match[2]))
	})
}

func part2(input string) int {
	r := regexp.MustCompile(`mul\(([0-9]*),([0-9]*)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	return utils.Reduce(r.FindAllStringIndex(input, -1), func(total int, matchIndexes []int, i int) int {
		precedingString := input[0:matchIndexes[0]]
		reversed := utils.ReverseString(precedingString)
		dont := strings.Index(reversed, ")(t'nod")
		do := strings.Index(reversed, ")(od")
		if dont == do || (dont > do && do != -1) {
			match := matches[i]
			return total + (utils.StringToInteger(match[1]) * utils.StringToInteger(match[2]))
		}
		return total
	})
}
