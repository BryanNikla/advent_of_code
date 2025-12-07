package registry

import (
	"sort"

	"advent_of_code/utils"
)

type SolutionRegistry map[int]map[int]func() utils.Solution

var Registry = make(SolutionRegistry)

func RegisterSolution(year int, day int, solution func() utils.Solution) {
	if _, exists := Registry[year]; !exists {
		Registry[year] = make(map[int]func() utils.Solution)
	}
	Registry[year][day] = solution
}

func (r SolutionRegistry) GetSolutions(year int, day int) []utils.Solution {
	daysMap, ok := r[year]
	if !ok {
		return nil
	}

	////////////////////////////////////////////
	// specific day
	if day > 0 {
		if solver, ok := daysMap[day]; ok {
			return []utils.Solution{solver()}
		}
		return nil
	}

	////////////////////////////////////////////
	// All days for the year
	var solutions []utils.Solution
	var days []int
	for d := range daysMap {
		days = append(days, d)
	}
	sort.Ints(days)
	for _, d := range days {
		solutions = append(solutions, daysMap[d]())
	}
	return solutions
}
