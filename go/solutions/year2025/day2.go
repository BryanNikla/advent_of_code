package year2025

import (
	"strconv"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 2, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 2)
		return utils.Solution{
			Day:   2,
			Test1: day2part1(input1) == 1227775554,
			Test2: day2part2(input2) == 4174379265,
		}
	})
}

func day2part1(input string) int {
	var sumOfInvalidIDs int
	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start := utils.StringToInteger(bounds[0])
		end := utils.StringToInteger(bounds[1])
		utils.ForRange(start, end, func(id int) {
			if isValueTwoSimilarParts(id) {
				sumOfInvalidIDs += id
			}
		})
	}
	return sumOfInvalidIDs
}

func day2part2(input string) int {
	var sumOfInvalidIDs int
	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start := utils.StringToInteger(bounds[0])
		end := utils.StringToInteger(bounds[1])
		utils.ForRange(start, end, func(id int) {
			if isValueSequenceOfRepeatedNumbers(id) {
				sumOfInvalidIDs += id
			}
		})
	}
	return sumOfInvalidIDs
}

func isValueTwoSimilarParts(value int) bool {
	valueStr := strconv.Itoa(value)
	length := len(valueStr)
	if length%2 == 0 {
		middle := length / 2
		return valueStr[:middle] == valueStr[middle:]
	}
	return false
}

func isValueSequenceOfRepeatedNumbers(value int) bool {
	valueStr := strconv.Itoa(value)
	length := len(valueStr)
	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			pattern := valueStr[:i]
			repeats := length / i
			strIfRepeated := strings.Repeat(pattern, repeats)
			if strIfRepeated == valueStr {
				return true
			}
		}
	}
	return false
}
