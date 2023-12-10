package day9

import (
	"os"
	"strconv"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

func allZeroes(seq []string) bool {
	for _, s := range seq {
		if s != "0" {
			return false
		}
	}
	return true
}

func nextElem(seq []string) int {
	var nextSeq []string
	for !allZeroes(seq) {
		for i := 1; i < len(seq); i++ {
			curr, _ := strconv.Atoi(seq[i])
			before, _ := strconv.Atoi(seq[i-1])
			nextSeq = append(nextSeq, strconv.Itoa(curr-before))
		}
		nextSeq = append(nextSeq, strconv.Itoa(nextElem(nextSeq)))
		lastNextElement, _ := strconv.Atoi(nextSeq[len(nextSeq)-1])
		lastCurrElement, _ := strconv.Atoi(seq[len(seq)-1])
		return lastNextElement + lastCurrElement
	}
	return 0
}

// Part1 returns the answer to Day 9, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 9, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	sum := 0

	for _, line := range lines {
		sum += nextElem(strings.Fields(line))
	}

	return sum
}

// Part2 returns the answer to Day 9, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 9, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	sum := 0

	for _, line := range lines {
		sum += nextElem(reverse(strings.Fields(line)))
	}

	return sum
}

func reverse(a []string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
