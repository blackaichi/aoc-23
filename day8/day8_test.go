package day8_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day8"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result1 := day8.Part1("./examplePart1-1.txt")
	result2 := day8.Part1("./examplePart1-2.txt")

	assert.Equal(t, 2, result1)
	assert.Equal(t, 6, result2)
}

func TestB(t *testing.T) {
	result := day8.Part2("./examplePart2.txt")

	assert.Equal(t, 6, result)
}
