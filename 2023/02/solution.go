package year2023

import (
	"regexp"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func Day2() utils.Solution {
	test1, test2 := utils.GetInputs(2023, 2)
	return utils.Solution{
		Day:   2,
		Test1: part1(test1) == 8,
		Test2: part2(test2) == 2286,
	}
}

func part1(input string) int {
	var sum int

	available := set{
		red:   12,
		green: 13,
		blue:  14,
	}

	for _, game := range inputToGames(input) {
		if utils.Every(game.sets, func(set set) bool {
			return setIsLegal(set, available)
		}) {
			sum = sum + game.id
		}
	}

	return sum
}

func part2(input string) int {
	var sum int
	for _, game := range inputToGames(input) {
		minSet := set{0, 0, 0}
		for _, set := range game.sets {
			if set.red > minSet.red {
				minSet.red = set.red
			}
			if set.green > minSet.green {
				minSet.green = set.green
			}
			if set.blue > minSet.blue {
				minSet.blue = set.blue
			}
		}
		sum = sum + (minSet.red * minSet.green * minSet.blue)
	}

	return sum
}

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}

func inputToGames(input string) []game {
	var games []game

	lines := utils.GetLines(input)

	for _, line := range lines {
		parts := strings.Split(line, ":")

		var game = game{
			id:   digitsOnlyToInt(parts[0]),
			sets: []set{},
		}

		for _, subset := range strings.Split(parts[1], ";") {
			var set = set{0, 0, 0}
			for _, pull := range strings.Split(subset, ",") {
				if strings.Contains(pull, "red") {
					set.red = digitsOnlyToInt(pull) + set.red
				}
				if strings.Contains(pull, "green") {
					set.green = digitsOnlyToInt(pull) + set.green
				}
				if strings.Contains(pull, "blue") {
					set.blue = digitsOnlyToInt(pull) + set.blue
				}
			}

			game.sets = append(game.sets, set)
		}

		games = append(games, game)
	}

	return games
}

func digitsOnlyToInt(str string) int {
	nonDigit := regexp.MustCompile(`[^0-9]`)
	x, _ := strconv.Atoi(nonDigit.ReplaceAllString(str, ""))
	return x
}

func setIsLegal(set set, available set) bool {
	return set.red <= available.red && set.green <= available.green && set.blue <= available.blue
}
