package year2025

import (
	"fmt"
	"strings"

	"advent_of_code/registry"
	"advent_of_code/utils"
)

func init() {
	registry.RegisterSolution(2025, 12, func() utils.Solution {
		input1, _ := utils.GetInput(2025, 12)
		return utils.Solution{
			Day:   12,
			Test1: day12part1(input1) == 2,
			Test2: false,
		}
	})
}

type RegionsToPopulate struct {
	X        int
	Y        int
	Presents map[int]int // shape index to count of how many must fit this region
}

type Present struct {
	Matrix [][]byte
	Area   int
}

func day12part1(input string) int {
	inputSegments := strings.Split(input, "\n\n")

	shapes := make(map[int][][]byte)

	// Parse all but the last segment as shape configurations
	for i, segment := range inputSegments[:len(inputSegments)-1] {
		lines := utils.GetLines(segment)
		shapes[i] = utils.LinesToByteMatrix(lines[1:])
	}

	presents := make([]Present, 0, len(shapes))
	for i, shape := range shapes {
		var area int
		utils.EachMatrix(shape, func(val byte, _ utils.Coordinates, _ [][]byte) {
			if val == '#' {
				area++
			}
		})
		presents[i] = Present{
			Matrix: shape,
			Area:   area,
		}
	}

	// last segment is regions to populate. Parse that also.
	regionLines := utils.GetLines(inputSegments[len(inputSegments)-1])

	regions := make([]RegionsToPopulate, 0, len(regionLines))

	for _, regionLine := range regionLines {
		parts := strings.Split(regionLine, ":")

		coords := strings.Split(parts[0], "x")

		presents := make(map[int]int)
		presentStrs := strings.Fields(parts[1])
		for i, presentStr := range presentStrs {
			presents[i] = utils.StringToInteger(presentStr)
		}

		regions = append(regions, RegionsToPopulate{
			X:        utils.StringToInteger(coords[0]),
			Y:        utils.StringToInteger(coords[1]),
			Presents: presents,
		})
	}

	/////////////////////////////////////////////////
	// Solve

	var regionsThatCanFitTheirPresents int

	for _, region := range regions {
		regionArea := region.X * region.Y
		var presentAreaTotal int
		for presentIndex, presentCount := range region.Presents {
			presentAreaTotal += presents[presentIndex].Area * presentCount
		}
		if regionArea <= presentAreaTotal {
			regionsThatCanFitTheirPresents++
		}
	}

	fmt.Printf("regionsThatCanFitTheirPresents: %d\n", regionsThatCanFitTheirPresents)
	return regionsThatCanFitTheirPresents
}
