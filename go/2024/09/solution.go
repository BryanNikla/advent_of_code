package year2024day09

import (
	"slices"
	"strings"

	utils "advent_of_code/utils"
)

func Solve() utils.Solution {
	input1, input2 := utils.GetInputs(2024, 9)
	return utils.Solution{
		Day:   9,
		Test1: part1(input1) == 1928,
		Test2: part2(input2) == 2858,
	}
}

func part1(input string) int {
	disk := prepDisk(strings.Split(input, ""))
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			swapIdx := slices.Index(disk, FREE_SPACE)
			if swapIdx != -1 {
				disk[i], disk[swapIdx] = disk[swapIdx], disk[i]
			}
		}
	}
	cleanDisk := make([]int, 0, len(disk))
	for _, v := range disk {
		if v != FREE_SPACE {
			cleanDisk = append(cleanDisk, v)
		}
	}
	return calculateChecksum(cleanDisk)
}

func part2(input string) int {
	disk := prepDisk(strings.Split(input, ""))

	var currentID = -10 // -10 is just an impossible value used as a reset
	var size int
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == currentID {
			size++
		} else {
			if size > 0 {
				// Do the thing
				i1, i2 := utils.FindFirstSpanWithLength(disk, -1, size)
				if i1 > -1 && i2 > -1 && i1 < i {
					utils.ReplaceValues(disk, currentID, FREE_SPACE)
					for _, idx := range utils.GetRange(i1, i2) {
						disk[idx] = currentID
					}
				}
				// Reset
				size = 0
				currentID = -10
			}
			// If this value is not FREE_SPACE, set it as our current w/ size = 1
			if disk[i] != FREE_SPACE {
				currentID = disk[i]
				size = 1
			}
		}
	}
	return calculateChecksum(disk)
}

var FREE_SPACE = -1

func prepDisk(input []string) []int {
	var disk []int
	var currentID int
	for idx, digit := range input {
		d := utils.StringToInteger(digit)
		for i := 0; i < d; i++ {
			if idx%2 == 0 {
				disk = append(disk, currentID)
			} else {
				disk = append(disk, FREE_SPACE)
			}
		}
		if idx%2 == 0 {
			currentID++
		}
	}
	return disk
}

func calculateChecksum(disk []int) int {
	var checksum int
	for i, v := range disk {
		if v != FREE_SPACE {
			checksum = checksum + (i * v)
		}
	}
	return checksum
}
