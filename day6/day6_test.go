package day6_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day6"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day6.Part1("./examplePart1.txt")

	assert.Equal(t, 288, result)
}

func TestB(t *testing.T) {
	result := day6.Part2("./examplePart1.txt")

	assert.Equal(t, 71503, result)
}
