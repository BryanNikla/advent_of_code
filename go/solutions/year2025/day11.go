package year2025

import (
	"fmt"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 11, func() utils.Solution {
		input1, input2 := utils.GetInput(2025, 11)
		return utils.Solution{
			Day:   11,
			Test1: day11part1(input1) == 5,
			Test2: day11part2(input2) == 2,
		}
	})
}

func day11part1(input string) int {
	lines := utils.GetLines(input)
	deviceMap := linesToDeviceMap(lines)

	cachedResults := make(map[string]int) // how many 'paths' from this device [key] to "out"

	var process func(device string) int
	process = func(device string) int {
		if device == "out" {
			return 1
		}
		if val, ok := cachedResults[device]; ok {
			return val
		}
		totalPaths := 0
		for _, conn := range deviceMap[device] {
			totalPaths += process(conn)
		}
		cachedResults[device] = totalPaths
		return totalPaths
	}

	return process("you")
}

func day11part2(input string) int {
	lines := utils.GetLines(input)
	deviceMap := linesToDeviceMap(lines)

	type CacheKey struct {
		device string
		hasDac bool
		hasFft bool
	}

	cachedResults := make(map[CacheKey]int)

	var process func(device string, hasDac bool, hasFft bool) int
	process = func(device string, hasDac bool, hasFft bool) int {
		if device == "dac" {
			hasDac = true
		}
		if device == "fft" {
			hasFft = true
		}

		if device == "out" {
			if hasDac && hasFft {
				return 1
			}
			return 0
		}

		key := CacheKey{device, hasDac, hasFft}

		// Already found, return the cached result
		if val, ok := cachedResults[key]; ok {
			return val
		}

		totalPaths := 0
		for _, conn := range deviceMap[device] {
			totalPaths += process(conn, hasDac, hasFft)
		}

		cachedResults[key] = totalPaths
		return totalPaths
	}

	answer := process("svr", false, false)
	fmt.Printf("ANSWER: %d\n", answer)
	return answer
}

func linesToDeviceMap(lines []string) map[string][]string {
	deviceMap := make(map[string][]string)
	for _, line := range lines {
		x := strings.Split(line, ":")
		device := strings.TrimSpace(x[0])
		connections := strings.Fields(strings.TrimSpace(x[1]))
		deviceMap[device] = connections
	}
	return deviceMap
}
