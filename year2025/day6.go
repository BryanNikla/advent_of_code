package year2025

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func SolutionDay6() utils.Solution {
	input := getInput(6)
	return utils.Solution{
		Day:   6,
		Test1: day6part1(input) == 4277556,
		//Test2: day5part2(input) == 14,
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func day6part1(input string) int {
	grandTotal := 0

	problems := make(map[int][]int)
	operations := make(map[int]string)

	for _, line := range utils.GetLines(input) {
		parts := strings.Fields(line)
		for i, part := range parts {
			if isNumeric(part) {
				problems[i] = append(problems[i], utils.StringToInteger(part))
			} else {
				operations[i] = part
			}
		}
	}

	//for _, problem := range problems {
	//	fmt.Printf("problem: %+v\n", problem)
	//}
	//for _, operation := range operations {
	//	fmt.Printf("operation: %+v\n", operation)
	//}

	for i := 0; i < len(problems); i++ {
		total := 0
		for j, num := range problems[i] {
			if j == 0 {
				total = num
			} else {
				op := operations[i]
				switch op {
				case "+":
					total += num
				case "*":
					total *= num
				}
			}
		}
		grandTotal += total
	}

	fmt.Printf("Grand Total: %d\n", grandTotal)
	return grandTotal
}
