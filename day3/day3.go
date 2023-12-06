package day3

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	timer "github.com/blackaichi/aoc-23"
)

func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func getAdjacentNumber(line string, index int) (int, string) {
	var adjacentNumber string
	r := []rune(line)
	for i := index; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			adjacentNumber = string(line[i]) + adjacentNumber
			r[i] = rune('.')
		} else {
			break
		}
	}
	for i := index + 1; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			adjacentNumber += string(line[i])
			r[i] = rune('.')
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(adjacentNumber)
	return num, string(r)
}

func getAdjacentNumbers(line string, symbols []int) ([]int, string) {
	var adjacentNumbers []int
	var adjacentNumber int
	if len(line) == 0 || len(symbols) == 0 {
		return adjacentNumbers, line
	}
	for _, symbol := range symbols {
		if symbol-1 < 0 || symbol-1 >= len(line) || symbol+1 >= len(line) {
			continue
		}
		if unicode.IsDigit(rune(line[symbol-1])) {
			adjacentNumber, line = getAdjacentNumber(line, symbol-1)
			adjacentNumbers = append(adjacentNumbers, adjacentNumber)
		}
		if unicode.IsDigit(rune(line[symbol+1])) {
			adjacentNumber, line = getAdjacentNumber(line, symbol+1)
			adjacentNumbers = append(adjacentNumbers, adjacentNumber)
		}
		if unicode.IsDigit(rune(line[symbol])) {
			adjacentNumber, line = getAdjacentNumber(line, symbol)
			adjacentNumbers = append(adjacentNumbers, adjacentNumber)
		}
	}

	return adjacentNumbers, line
}

func getSymbolIndices(line string) []int {
	indices := []int{}
	for pos, charAsRune := range line {
		if charAsRune != 46 && (charAsRune < 48 || charAsRune > 57) {
			indices = append(indices, pos)
		}
	}
	return indices
}

// Part1 returns the answer to Day 2, Part 1 of the Advent
// of Code challenge 2023.
func Part1() int {
	defer timer.Timer("Day 3, Part 1")()
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	var sum int
	var symbols []int
	var ansSymbols []int
	var ansLine string

	for _, line := range strings.Split(string(input), "\n") {
		symbols = getSymbolIndices(line)

		numbers, _ := getAdjacentNumbers(ansLine, symbols)
		sum += sumArray(numbers)
		numbers, _ = getAdjacentNumbers(line, ansSymbols)
		sum += sumArray(numbers)
		numbers, ansLine = getAdjacentNumbers(line, symbols)
		sum += sumArray(numbers)

		ansSymbols = symbols
	}

	return sum
}

// Part2 returns the answer to Day 2, Part 2 of the Advent
// of Code challenge 2023.
func Part2() int {
	defer timer.Timer("Day 3, Part 2")()
	input, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	var sum int
	var numbers []int

	lines := strings.Split(string(input), "\n")
	for i := 0; i < len(lines); i++ {
		regex := regexp.MustCompile(`[*]`)
		match := regex.FindAllStringIndex(lines[i], -1)

		for _, m := range match {
			numbers = []int{}
			num, _ := getAdjacentNumbers(lines[i-1], []int{m[0]})
			numbers = append(numbers, num...)
			num, _ = getAdjacentNumbers(lines[i], []int{m[0]})
			numbers = append(numbers, num...)
			num, _ = getAdjacentNumbers(lines[i+1], []int{m[0]})
			numbers = append(numbers, num...)

			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			}
		}
	}

	return sum
}
