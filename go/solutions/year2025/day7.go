package year2025

import (
	"fmt"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 7, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 7)
		return utils.Solution{
			Day:   7,
			Test1: day7part1(input1) == 21,
			Test2: day7part2(input2) == 40,
		}
	})
}

const (
	TACHYON_BEAM   = '|'
	BEAM_SPLITTER  = '^'
	START_POSITION = 'S'
)

func day7part1(input string) int {
	lines := utils.GetLines(input)
	beamsAt := map[int]bool{} // Track where the beams are currently
	totalSplits := 0
	for i := range lines {
		// Find all index of beams
		for ci, r := range lines[i] {
			if r == START_POSITION {
				beamsAt[ci] = true
				continue
			}
			if r == TACHYON_BEAM {
				beamsAt[ci] = true
			}
			if r == BEAM_SPLITTER {
				if beamsAt[ci] {
					totalSplits++
					beamsAt[ci] = false
					beamsAt[ci-1] = true
					beamsAt[ci+1] = true
				}
			}
		}
	}
	return totalSplits
}

func day7part2(input string) int {
	lines := utils.GetLines(input)

	// Using a cache for results. Without this the recursive process would run for an actual eternity..
	// Key = "index,depth"
	// Value = count of paths from here (x,y)
	cachedResults := make(map[string]int)

	var processBeam func(beamAt int, depth int) int
	processBeam = func(beamAtIndex int, depth int) int {
		// Bottom of input, return 1
		if depth == len(lines) {
			return 1
		}

		// Create cachedResults key
		key := fmt.Sprintf("%d,%d", beamAtIndex, depth)

		// If value from this position was already found, just return that
		if val, foundAlready := cachedResults[key]; foundAlready {
			return val
		}

		totalPaths := 0
		if lines[depth][beamAtIndex] == BEAM_SPLITTER {
			// Split the beam, and add results from two processes
			totalPaths += processBeam(beamAtIndex-1, depth+1)
			totalPaths += processBeam(beamAtIndex+1, depth+1)
		} else {
			// No split, continue in place
			totalPaths += processBeam(beamAtIndex, depth+1)
		}

		cachedResults[key] = totalPaths
		return totalPaths
	}

	// Find the beam starting position
	startIndex := 0
	for ci, r := range lines[0] {
		if r == START_POSITION {
			startIndex = ci
			break

		}
	}

	return processBeam(startIndex, 0)
}
