package day3_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day3"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day3.Part1("./examplePart1.txt")

	assert.Equal(t, 4361, result)
}

func TestB(t *testing.T) {
	result := day3.Part2("./examplePart1.txt")

	assert.Equal(t, 467835, result)
}
