package day11_test

import (
	"testing"

	"github.com/blackaichi/aoc-23/day11"
	"github.com/magiconair/properties/assert"
)

func TestA(t *testing.T) {
	result := day11.Part1("./examplePart1.txt")

	assert.Equal(t, 374, result)
}
