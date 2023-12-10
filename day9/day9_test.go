package day9_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day9"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day9.Part1("./examplePart1.txt")

	assert.Equal(t, 114, result)
}

func TestB(t *testing.T) {
	result := day9.Part2("./examplePart1.txt")

	assert.Equal(t, 2, result)
}
