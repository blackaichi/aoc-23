package day2_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day2"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day2.Part1("./examplePart1.txt")

	assert.Equal(t, 8, result)
}

func TestB(t *testing.T) {
	result := day2.Part2("./examplePart1.txt")

	assert.Equal(t, 2286, result)
}
