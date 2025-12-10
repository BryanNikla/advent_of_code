package year2025

import (
	"fmt"
	"regexp"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 10, func() utils.Solution {
		input1, _ := utils.GetInput(2025, 10)
		return utils.Solution{
			Day:   10,
			Test1: day10part1(input1) == 7,
			Test2: false,
		}
	})
}

func day10part1(input string) int {
	lines := utils.GetLines(input)

	machineConfigs := utils.Reduce(lines, func(acc []MachineConfig, line string, _ int) []MachineConfig {
		return append(acc, parseLineToMachineConfig(line))
	})

	fmt.Println(machineConfigs[0])

	// TODO: Now actually solve the problem :|

	return 0
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
