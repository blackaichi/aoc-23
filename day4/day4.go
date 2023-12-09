package day4

import (
	"os"
	"slices"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

func splitCardsAndWinNum(line string) ([]string, []string) {
	cards := strings.Split(strings.Split(line, ":")[1], "|")
	winNum := strings.Fields(cards[0])
	numHand := strings.Fields(cards[1])
	return winNum, numHand
}

// Part1 returns the answer to Day 4, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 4, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var sum int
	var points int

	for _, line := range strings.Split(string(input), "\n") {
		points = 0

		winNum, numHand := splitCardsAndWinNum(line)

		for _, num := range numHand {
			if slices.Contains(winNum, num) {
				points++
			}
		}
		if points > 0 {
			sum += 1 << (points - 1)
		}

	}

	return sum
}

// Part2 returns the answer to Day 4, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 4, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var points int

	lines := strings.Split(string(input), "\n")
	cardCopies := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cardCopies[i] = 1
	}

	for i, line := range lines {
		points = 0

		winNum, numHand := splitCardsAndWinNum(line)

		for _, num := range numHand {
			if slices.Contains(winNum, num) {
				points++
			}
		}

		for j := i + 1; j < len(lines) && points > 0; j++ {
			cardCopies[j] += cardCopies[i]
			points--
		}
	}

	sum := sumArray(cardCopies)
	return sum
}

func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
