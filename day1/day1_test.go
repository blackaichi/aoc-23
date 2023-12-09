package day1_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day1"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day1.Part1("./examplePart1.txt")

	assert.Equal(t, 142, result)
}

func TestB(t *testing.T) {
	result := day1.Part2("./examplePart2.txt")

	assert.Equal(t, 281, result)
}
