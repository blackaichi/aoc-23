package day2

import (
	"os"
	"strconv"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

// Part1 returns the answer to Day 2, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 1, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	var sum int
	for _, line := range strings.Split(string(input), "\n") {
		correctLine := true
		res1 := strings.Split(line, ":")
		res2 := strings.Split(res1[1], ";")
		for i := 0; i < len(res2) && correctLine; i++ {
			res2[i] = strings.TrimSpace(res2[i])
			res3 := strings.Split(res2[i], ",")
			currRed := maxRed
			currGreen := maxGreen
			currBlue := maxBlue
			for j := 0; j < len(res3); j++ {
				res3[j] = strings.TrimSpace(res3[j])
				res4 := strings.Split(res3[j], " ")
				value, _ := strconv.Atoi(res4[0])
				if res4[1] == "red" {
					currRed -= value
				} else if res4[1] == "green" {
					currGreen -= value
				} else if res4[1] == "blue" {
					currBlue -= value
				} else {
					panic("Invalid color")
				}
			}

			if currRed < 0 || currGreen < 0 || currBlue < 0 {
				correctLine = false
				break
			}
		}

		if correctLine {
			game, _ := strconv.Atoi(strings.Split(res1[0], " ")[1])
			sum += game
		}
	}

	return sum
}

// Part2 returns the answer to Day 2, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 1, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, line := range strings.Split(string(input), "\n") {
		min_red := 0
		min_green := 0
		min_blue := 0
		res1 := strings.Split(line, ":")
		clean_line := strings.ReplaceAll(res1[1], ";", ",")
		res2 := strings.Split(clean_line, ",")
		for i := 0; i < len(res2); i++ {
			res2[i] = strings.TrimSpace(res2[i])
			res3 := strings.Split(res2[i], " ")
			value, _ := strconv.Atoi(res3[0])
			if res3[1] == "red" && value > min_red {
				min_red = value
			} else if res3[1] == "green" && value > min_green {
				min_green = value
			} else if res3[1] == "blue" && value > min_blue {
				min_blue = value
			}
		}

		sum += min_red * min_green * min_blue
	}

	return sum
}
