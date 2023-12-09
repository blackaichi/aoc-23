package day4_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day4"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day4.Part1("./examplePart1.txt")

	assert.Equal(t, 13, result)
}

func TestB(t *testing.T) {
	result := day4.Part2("./examplePart1.txt")

	assert.Equal(t, 30, result)
}
