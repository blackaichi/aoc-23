package day6

import (
	"os"
	"strconv"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

// Part1 returns the answer to Day 6, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 6, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	total := 1

	lines := strings.Split(string(input), "\n")
	time := strings.Fields(strings.Split(lines[0], ":")[1])
	distance := strings.Fields(strings.Split(lines[1], ":")[1])
	for i := 0; i < len(time); i++ {
		sum := 0
		t, _ := strconv.Atoi(time[i])
		d, _ := strconv.Atoi(distance[i])
		for j := 0; j < t; j++ {
			if j*(t-j) > d {
				sum++
			}
		}
		total = total * sum
	}

	return total
}

// Part2 returns the answer to Day 6, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 6, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var minWin, maxWIin int

	lines := strings.Split(string(input), "\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			minWin = i
			break
		}
	}
	for i := time; i > 0; i-- {
		if i*(time-i) > distance {
			maxWIin = i
			break
		}
	}

	return maxWIin - minWin + 1
}
