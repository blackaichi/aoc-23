package day10_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day10"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result1 := day10.Part1("./examplePart1-1.txt")
	result2 := day10.Part1("./examplePart1-2.txt")

	assert.Equal(t, 4, result1)
	assert.Equal(t, 8, result2)
}

func TestB(t *testing.T) {
	result1 := day10.Part2("./examplePart2-1.txt")
	result2 := day10.Part2("./examplePart2-2.txt")
	result3 := day10.Part2("./examplePart2-3.txt")

	assert.Equal(t, 4, result1)
	assert.Equal(t, 8, result2)
	assert.Equal(t, 10, result3)
}
