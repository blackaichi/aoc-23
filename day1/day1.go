package day1

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

func checkNumber(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		return false
	}
	return true
}

// Part1 returns the answer to Day 1, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 1, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var sum int
	var buffer bytes.Buffer
	for _, line := range strings.Split(string(input), "\n") {
		buffer = bytes.Buffer{}
		for j := 0; j < len(line); j++ {
			if checkNumber(string(line[j])) {
				buffer.WriteByte(line[j])
				break
			}
		}

		for j := len(line) - 1; j >= 0; j-- {
			if checkNumber(string(line[j])) {
				buffer.WriteByte(line[j])
				break
			}
		}
		num, err := strconv.Atoi(buffer.String())
		if err != nil {
			panic(err)
		}
		sum += num
	}

	return sum
}

// Part2 returns the answer to Day 1, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 1, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, line := range strings.Split(string(input), "\n") {
		firstNumber := findFirstNumber(line)
		secondNumber := findLastNumber(line)

		num := firstNumber*10 + secondNumber
		if err != nil {
			panic(err)
		}
		sum += num
	}

	return sum
}

func findFirstNumber(s string) int {
	regex := regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	matches := regex.FindString(s)
	if matches != "" {
		if _, err := strconv.Atoi(matches); err == nil {
			res, _ := string_to_int[matches[0:1]]
			return res
		}
		return string_to_int[matches]
	}
	return -1
}

func findLastNumber(s string) int {
	regex := regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	for i := len(s) - 1; i >= 0; i-- {
		match := regex.FindString(string(s[i:]))
		if match != "" {
			if _, err := strconv.Atoi(match); err == nil {
				return string_to_int[match[len(match)-1:]]
			}
			return string_to_int[match]
		}
	}
	return -1
}

var string_to_int = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
