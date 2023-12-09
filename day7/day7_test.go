package day7_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day7"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day7.Part1("./examplePart1.txt")

	assert.Equal(t, 6440, result)
}

func TestB(t *testing.T) {
	result := day7.Part2("./examplePart1.txt")

	assert.Equal(t, 5905, result)
}
