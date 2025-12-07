package year2025

import (
	"fmt"

	"advent_of_code/utils"
)

func getInput(day int) string {
	var path = fmt.Sprintf("year2025/inputs/day%d.txt", day)
	return utils.GetFileContent(path)
}
