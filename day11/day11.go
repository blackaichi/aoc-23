package day11

import (
	"os"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

type coord struct {
	x int
	y int
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func getDistanceBetweenGalaxies(g1, g2 coord, expSpaceX, expSpaceY []int) int {
	totalDistance := 0
	gi := g1
	gj := g2
	if g1.x > g2.x {
		gi, gj = g2, g1
	}
	for k := gi.x + 1; k <= gj.x; k++ {
		if contains(expSpaceX, k) {
			totalDistance++
		} else {
			totalDistance += 2
		}
	}
	for k := g1.y + 1; k <= g2.y; k++ {
		if contains(expSpaceY, k) {
			totalDistance++
		} else {
			totalDistance += 2
		}
	}
	return totalDistance
}

// Part1 returns the answer to Day 11, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 11, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var galaxies []coord
	var containsGalaxyX, containsGalaxyY []int

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, coord{j, i})
				containsGalaxyY = append(containsGalaxyY, i)
				containsGalaxyX = append(containsGalaxyX, j)
			}
		}
	}

	totalDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += getDistanceBetweenGalaxies(galaxies[i], galaxies[j], containsGalaxyX, containsGalaxyY)
		}
	}

	return totalDistance
}

// Part2 returns the answer to Day 11, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 11, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var galaxies []coord
	var containsGalaxyX, containsGalaxyY []int

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, coord{j, i})
				containsGalaxyY = append(containsGalaxyY, i)
				containsGalaxyX = append(containsGalaxyX, j)
			}
		}
	}

	totalDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += getDistanceBetweenGalaxiesX100000(galaxies[i], galaxies[j], containsGalaxyX, containsGalaxyY)
		}
	}

	return totalDistance
}

func getDistanceBetweenGalaxiesX100000(g1, g2 coord, expSpaceX, expSpaceY []int) int {
	totalDistance := 0
	gi := g1
	gj := g2
	if g1.x > g2.x {
		gi, gj = g2, g1
	}
	for k := gi.x + 1; k <= gj.x; k++ {
		if contains(expSpaceX, k) {
			totalDistance++
		} else {
			totalDistance += 1000000
		}
	}
	for k := g1.y + 1; k <= g2.y; k++ {
		if contains(expSpaceY, k) {
			totalDistance++
		} else {
			totalDistance += 1000000
		}
	}
	return totalDistance
}
