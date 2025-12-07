package year2024

import (
	"regexp"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2024, 3, func() utils.Solution {
		testInput1, testInput2 := utils.GetInput(2024, 3)
		return utils.Solution{
			Day:   3,
			Test1: day3part1(testInput1) == 161,
			Test2: day3part2(testInput2) == 48,
		}
	})
}

func day3part1(input string) int {
	r := regexp.MustCompile(`mul\(([0-9]*),([0-9]*)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	return utils.Reduce(matches, func(total int, match []string, _ int) int {
		return total + (utils.StringToInteger(match[1]) * utils.StringToInteger(match[2]))
	})
}

func day3part2(input string) int {
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
