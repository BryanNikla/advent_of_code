package year2024day05

import (
	"slices"
	"sort"
	"strings"

	"advent_of_code/utils"
)

func Solve() utils.Solution {
	testInput1, testInput2 := utils.GetInputs(2024, 5)
	return utils.Solution{
		Day:   5,
		Test1: part1(testInput1) == 143,
		Test2: part2(testInput2) == 123,
	}
}

func part1(input string) int {
	lines := utils.GetLines(input)
	mustBeBefore, mustBeAfter, updates := prepare(lines)
	var total int
	for _, update := range updates {
		if verifyUpdate(update, mustBeBefore, mustBeAfter) {
			total = total + utils.MiddleItem(update)
		}
	}

	return total
}

func part2(input string) int {
	lines := utils.GetLines(input)
	mustBeBefore, mustBeAfter, updates := prepare(lines)
	var total int
	for _, update := range updates {
		if !verifyUpdate(update, mustBeBefore, mustBeAfter) {
			var x []Val
			for _, i := range update {
				x = append(x, Val{i, mustBeAfter})
			}
			sort.Sort(FixOrder(x))
			middle := utils.MiddleItem(x)
			total = total + middle.I
		}
	}

	return total
}

type Val struct {
	I            int
	MustBeBefore map[int][]int
}
type FixOrder []Val

func (a FixOrder) Len() int {
	return len(a)
}
func (a FixOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a FixOrder) Less(i, j int) bool {
	return slices.Contains(a[i].MustBeBefore[a[i].I], a[j].I)
}

func verifyUpdate(update []int, mustBeBefore map[int][]int, mustBeAfter map[int][]int) bool {
	for i, val := range update {
		for ii := 0; ii < i; ii++ {
			if slices.Contains(mustBeAfter[val], update[ii]) {
				return false
			}
		}
		for ii := i + 1; ii < len(update)-1; ii++ {
			if slices.Contains(mustBeBefore[val], update[ii]) {
				return false
			}
		}
	}
	return true
}

func prepare(lines []string) (map[int][]int, map[int][]int, [][]int) {
	mustBeBefore := make(map[int][]int)
	mustBeAfter := make(map[int][]int)

	var updates [][]int

	var parsingRules = true
	for _, line := range lines {
		if len(line) == 0 {
			parsingRules = false
			continue
		}

		if parsingRules {
			parts := strings.Split(line, "|")
			left := utils.StringToInteger(parts[0])
			right := utils.StringToInteger(parts[1])
			mustBeBefore[right] = append(mustBeBefore[right], left)
			mustBeAfter[left] = append(mustBeAfter[left], right)
		} else {
			var update []int
			for _, x := range strings.Split(line, ",") {
				update = append(update, utils.StringToInteger(x))
			}
			updates = append(updates, update)
		}
	}

	return mustBeBefore, mustBeAfter, updates
}
