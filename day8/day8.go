package day8

import (
	"os"
	"regexp"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

type pair [2]string

var nonAlphanumericRegex = regexp.MustCompile(`[^A-Z0-9\,]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

// Part1 returns the answer to Day 8, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 8, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	instructions := lines[0]
	nodes := make(map[string]pair)
	steps := 0

	for i := 2; i < len(lines); i++ {
		aux := strings.Split(lines[i], "=")
		key := strings.TrimSpace(aux[0])
		values := strings.Split(clearString(aux[1]), ",")
		nodes[key] = pair{values[0], values[1]}
	}

	current := "AAA"
	next := ""
	for i := 0; ; i = (i + 1) % len(instructions) {
		if instructions[i] == 'L' {
			next = nodes[current][0]
		} else {
			next = nodes[current][1]
		}
		current = next
		steps++
		if current == "ZZZ" {
			break
		}
	}

	return steps
}

// Part2 returns the answer to Day 8, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 8, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	instructions := lines[0]
	nodes := make(map[string]pair)
	steps := 0
	var startNodes []string

	for i := 2; i < len(lines); i++ {
		aux := strings.Split(lines[i], "=")
		key := strings.TrimSpace(aux[0])
		values := strings.Split(clearString(aux[1]), ",")
		nodes[key] = pair{values[0], values[1]}
		if key[2] == 'A' {
			startNodes = append(startNodes, key)
		}
	}
	var cycle []int

	for i := 0; ; i = (i + 1) % len(instructions) {
		for j := 0; j < len(startNodes); j++ {
			if instructions[i] == 'L' {
				startNodes[j] = nodes[startNodes[j]][0]
			} else {
				startNodes[j] = nodes[startNodes[j]][1]
			}
			if startNodes[j][2] == 'Z' {
				cycle = append(cycle, steps+1)
				startNodes = append(startNodes[:j], startNodes[j+1:]...)
				j--
			}
		}
		if len(startNodes) == 0 {
			break
		}
		steps++
	}
	result := LCM(cycle[0], cycle[1], cycle[2:]...)

	return result
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
