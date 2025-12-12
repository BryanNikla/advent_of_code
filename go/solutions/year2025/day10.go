package year2025

import (
	"fmt"
	"regexp"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 10, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 10)
		return utils.Solution{
			Day:   10,
			Test1: day10part1(input1) == 7,
			Test2: day10part2(input2) == 33,
		}
	})
}

func day10part1(input string) int {
	lines := utils.GetLines(input)

	machineConfigs := utils.Reduce(lines, func(acc []MachineConfig, line string, _ int) []MachineConfig {
		return append(acc, parseLineToMachineConfig(line))
	})

	return utils.Reduce(machineConfigs, func(acc int, config MachineConfig, _ int) int {
		clicks := solve(config.Buttons, config.Lights)
		return acc + clicks
	})
}

func day10part2(input string) int {
	lines := utils.GetLines(input)

	machineConfigs := utils.Reduce(lines, func(acc []MachineConfig, line string, _ int) []MachineConfig {
		return append(acc, parseLineToMachineConfig(line))
	})

	solution := utils.Reduce(machineConfigs, func(acc int, config MachineConfig, _ int) int {
		clicks := solve(config.Buttons, config.Lights)
		return acc + clicks
	})

	fmt.Printf("SOLUTION: %d \n", solution)
	return solution
}

type LightStatus byte

const (
	LightOff = '.'
	LightOn  = '#'
)

type MachineConfig struct {
	Lights     []LightStatus
	Buttons    [][]int
	JoltageReq []int // ignore for part 1
}

// parseLineToMachineConfig parses a line like this:
// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// into lights, buttons, and JoltageReq req
func parseLineToMachineConfig(line string) MachineConfig {
	reLights := regexp.MustCompile(fmt.Sprintf(`\[([%c%c]+)]`, LightOn, LightOff))
	reButtons := regexp.MustCompile(`\(([\d,]+)\)`)
	reJoltage := regexp.MustCompile(`\{([\d,]+)}`)

	var config MachineConfig
	if matches := reLights.FindStringSubmatch(line); len(matches) > 1 {
		rawLights := matches[1] // e.g., ".##."
		config.Lights = make([]LightStatus, len(rawLights))
		for i := 0; i < len(rawLights); i++ {
			config.Lights[i] = LightStatus(rawLights[i])
		}
	}
	buttonMatches := reButtons.FindAllStringSubmatch(line, -1)
	for _, match := range buttonMatches {
		nums := utils.ParseIntList(match[1])
		config.Buttons = append(config.Buttons, nums)
	}
	if matches := reJoltage.FindStringSubmatch(line); len(matches) > 1 {
		config.JoltageReq = utils.ParseIntList(matches[1])
	}
	return config
}

func buttonsToWiringMatrix(buttons [][]int, numLights int) [][]int {
	numButtons := len(buttons)

	matrix := make([][]int, numLights)
	for i := range matrix {
		matrix[i] = make([]int, numButtons)
	}

	for btnIndex, affectedLights := range buttons {
		for _, lightIndex := range affectedLights {
			if lightIndex < numLights {
				matrix[lightIndex][btnIndex] = 1
			}
		}
	}

	return matrix
}

// solve performs Gaussian Elimination and brute-forces free variables for the minimal solution
func solve(buttons [][]int, targetLights []LightStatus) int {
	totalLights := len(targetLights)

	// Build Target Vector
	target := make([]int, totalLights)
	for i, status := range targetLights {
		if status == LightOn {
			target[i] = 1
		}
	}

	// Build Matrix
	wiringMatrix := buttonsToWiringMatrix(buttons, totalLights)
	rowCount := len(wiringMatrix)
	colCount := len(wiringMatrix[0])

	// Build Augmented Matrix [A | b]
	aug := make([][]int, rowCount)
	for i := range wiringMatrix {
		aug[i] = make([]int, colCount+1)
		copy(aug[i], wiringMatrix[i])
		aug[i][colCount] = target[i]
	}

	// --- Gaussian Elimination (Gauss-Jordan) ---
	pivotRow := 0
	pivots := make([]int, rowCount) // maps row -> col of the pivot
	for i := range pivots {
		pivots[i] = -1
	}

	// colToPivotRow maps a column index to the row index where its pivot lives
	colToPivotRow := make(map[int]int)

	for col := 0; col < colCount && pivotRow < rowCount; col++ {
		selRow := -1
		for i := pivotRow; i < rowCount; i++ {
			if aug[i][col] == 1 {
				selRow = i
				break
			}
		}

		if selRow == -1 {
			continue // This is a FREE variable
		}

		aug[pivotRow], aug[selRow] = aug[selRow], aug[pivotRow]

		for i := 0; i < rowCount; i++ {
			if i != pivotRow && aug[i][col] == 1 {
				for j := col; j <= colCount; j++ {
					aug[i][j] ^= aug[pivotRow][j]
				}
			}
		}

		pivots[pivotRow] = col
		colToPivotRow[col] = pivotRow
		pivotRow++
	}

	// Check for impossible systems (0 = 1)
	for i := pivotRow; i < rowCount; i++ {
		if aug[i][colCount] == 1 {
			return 0 // No solution possible
		}
	}

	// Identify free variables
	var freeCols []int
	for col := 0; col < colCount; col++ {
		if _, isPivot := colToPivotRow[col]; !isPivot {
			freeCols = append(freeCols, col)
		}
	}

	minClicks := int(^uint(0) >> 1) // Max Int

	// Iterate 2^k combinations for k free variables
	countFree := len(freeCols)
	combinations := 1 << countFree

	for i := 0; i < combinations; i++ {
		candidateSol := make([]int, colCount)

		currentClicks := 0
		for bit := 0; bit < countFree; bit++ {
			if (i & (1 << bit)) != 0 {
				candidateSol[freeCols[bit]] = 1
				currentClicks++
			}
		}

		for r := 0; r < pivotRow; r++ {
			c := pivots[r]
			val := aug[r][colCount] // start with b vector

			// XOR with relevant free variables in this row
			for _, freeC := range freeCols {
				if aug[r][freeC] == 1 {
					val ^= candidateSol[freeC]
				}
			}
			candidateSol[c] = val
			if val == 1 {
				currentClicks++
			}
		}

		if currentClicks < minClicks {
			minClicks = currentClicks
		}
	}

	return minClicks
}
