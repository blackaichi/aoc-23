package day10

import (
	"math"
	"os"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

type coord struct {
	x int
	y int
}

var allowedLeft = []byte{'-', 'F', 'L', 'S'}
var allowedRight = []byte{'-', 'J', '7', 'S'}
var allowedUp = []byte{'|', 'F', '7', 'S'}
var allowedDown = []byte{'|', 'J', 'L', 'S'}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func checkLeft(line string, posX int) bool {
	if posX > 0 && contains(allowedLeft, line[posX-1]) && contains(allowedRight, line[posX]) {
		return true
	}
	return false
}

func checkRight(line string, posX int) bool {
	if posX < len(line)-1 && contains(allowedRight, line[posX+1]) && contains(allowedLeft, line[posX]) {
		return true
	}
	return false
}

func checkUp(line []string, posX, posY int) bool {
	if posY > 0 && contains(allowedUp, line[posY-1][posX]) && contains(allowedDown, line[posY][posX]) {
		return true
	}
	return false
}

func checkDown(line []string, posX, posY int) bool {
	if posY < len(line)-1 && contains(allowedDown, line[posY+1][posX]) && contains(allowedUp, line[posY][posX]) {
		return true
	}
	return false
}

func pathFinder(lines []string, posX, posY int, history []coord) int {
	if checkLeft(lines[posY], posX) && !contains(history, coord{posX - 1, posY}) {
		return pathFinder(lines, posX-1, posY, append(history, coord{posX, posY})) + 1
	}
	if checkRight(lines[posY], posX) && !contains(history, coord{posX + 1, posY}) {
		return pathFinder(lines, posX+1, posY, append(history, coord{posX, posY})) + 1
	}
	if checkUp(lines, posX, posY) && !contains(history, coord{posX, posY - 1}) {
		return pathFinder(lines, posX, posY-1, append(history, coord{posX, posY})) + 1
	}
	if checkDown(lines, posX, posY) && !contains(history, coord{posX, posY + 1}) {
		return pathFinder(lines, posX, posY+1, append(history, coord{posX, posY})) + 1
	}
	return 1
}

func findStart(lines []string) (int, int) {
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				return x, y
			}
		}
	}
	return 0, 0
}

// Part1 returns the answer to Day 10, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 10, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	posX, posY := findStart(lines)
	steps := pathFinder(lines, posX, posY, []coord{})

	return steps / 2
}

// Part2 returns the answer to Day 10, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 10, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	posX, posY := findStart(lines)
	mainLoop := mainLoopFinder(lines, posX, posY, []coord{})
	dots := findCandidateEnclose(lines, mainLoop)

	var enclosedDots []coord
	for _, dot := range dots {
		if !isOutside(lines, mainLoop, dot) {
			enclosedDots = append(enclosedDots, dot)
		}
	}

	return len(enclosedDots)
}

func isOutside(lines []string, mainLoop []coord, candidate coord) bool {
	nLines := 0
	sum1 := []byte{'F', 'J', '-', '|'}
	x, y := candidate.x-1, candidate.y-1
	for x >= 0 && y >= 0 {
		b := lines[y][x]
		if b == 'S' {
			b = transformS(mainLoop)
		}
		if contains(sum1, b) && isMainLoop(x, y, mainLoop) {
			nLines++
		}
		x--
		y--
	}

	return nLines%2 == 0
}

func transformS(mainLoop []coord) byte {
	beforeS := mainLoop[len(mainLoop)-1]
	afterS := mainLoop[1]
	if (math.Abs(float64(beforeS.y-afterS.y)) == 2) || (math.Abs(float64(beforeS.x-afterS.x)) == 2) {
		return '-'
	} else if (beforeS.y == afterS.y+1 && beforeS.x+1 == afterS.x) || (beforeS.x == afterS.x+1 && beforeS.y+1 == afterS.y) {
		return 'J'
	}
	return 'L'
}

func findCandidateEnclose(lines []string, mainLoop []coord) []coord {
	dots := []coord{}
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if !isMainLoop(j, i, mainLoop) {
				dots = append(dots, coord{j, i})
			}
		}
	}
	return dots
}

func isMainLoop(posX, posY int, mainLoop []coord) bool {
	if contains(mainLoop, coord{posX, posY}) {
		return true
	}
	return false
}

func mainLoopFinder(lines []string, posX, posY int, history []coord) []coord {
	if checkLeft(lines[posY], posX) && !contains(history, coord{posX - 1, posY}) {
		return mainLoopFinder(lines, posX-1, posY, append(history, coord{posX, posY}))
	}
	if checkRight(lines[posY], posX) && !contains(history, coord{posX + 1, posY}) {
		return mainLoopFinder(lines, posX+1, posY, append(history, coord{posX, posY}))
	}
	if checkUp(lines, posX, posY) && !contains(history, coord{posX, posY - 1}) {
		return mainLoopFinder(lines, posX, posY-1, append(history, coord{posX, posY}))
	}
	if checkDown(lines, posX, posY) && !contains(history, coord{posX, posY + 1}) {
		return mainLoopFinder(lines, posX, posY+1, append(history, coord{posX, posY}))
	}
	return append(history, coord{posX, posY})
}
