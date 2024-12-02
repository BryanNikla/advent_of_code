package go2024

import (
	"advent_of_code/utils"
)

func Execute(day int) utils.Solution {
	switch day {
	case 1:
		return Day1()
	case 2:
		return Day2()
	}

	return utils.Solution{}
}
