package day7

import (
	"bytes"
	"os"
	"sort"
	"strconv"
	"strings"

	timer "github.com/blackaichi/aoc-23"
)

type hand int

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type player struct {
	cards string
	bid   int
}

func (p *player) getHand() hand {
	if strings.Count(p.cards, string(p.cards[0])) == 5 {
		return FiveOfAKind
	}
	for i := 0; i < len(p.cards); i++ {
		if strings.Count(p.cards, string(p.cards[i])) == 4 {
			return FourOfAKind
		}
		if strings.Count(p.cards, string(p.cards[i])) == 3 {
			for j := i + 1; j < len(p.cards); j++ {
				if strings.Count(p.cards, string(p.cards[j])) == 2 {
					return FullHouse
				}
			}
			return ThreeOfAKind
		}
		if strings.Count(p.cards, string(p.cards[i])) == 2 {
			for j := i + 1; j < len(p.cards); j++ {
				if strings.Count(p.cards, string(p.cards[j])) == 2 && p.cards[i] != p.cards[j] {
					return TwoPair
				} else if strings.Count(p.cards, string(p.cards[j])) == 3 {
					return FullHouse
				}
			}
			return OnePair
		}
	}
	return HighCard
}

func getPlayers(lines []string) []player {
	var players []player
	for i := 0; i < len(lines); i++ {
		cards := strings.Fields(lines[i])[0]
		bid, _ := strconv.Atoi(strings.Fields(lines[i])[1])
		p := player{cards, bid}
		players = append(players, p)
	}
	return players
}

func greaterCard(a, b rune) bool {
	order := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	aIndex := bytes.IndexRune(order, a)
	bIndex := bytes.IndexRune(order, b)
	return aIndex > bIndex
}

func orderCards(p []player) []player {
	sort.Slice(p, func(i, j int) bool {
		for k := 0; k < len(p[i].cards); k++ {
			if p[i].cards[k] != p[j].cards[k] {
				return greaterCard(rune(p[i].cards[k]), rune(p[j].cards[k]))
			}
		}
		return false
	})
	return p
}

func getWinnings(p []player, ranking int) int {
	sum := 0
	for i := 0; i < len(p); i++ {
		sum += p[i].bid * (ranking + i)
	}
	return sum
}

// Part1 returns the answer to Day 7, Part 1 of the Advent
// of Code challenge 2023.
func Part1(filePath string) int {
	defer timer.Timer("Day 7, Part 1")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var totalWinnings int
	var playerByHand [7][]player

	lines := strings.Split(string(input), "\n")

	players := getPlayers(lines)

	for i := 0; i < len(players); i++ {
		switch players[i].getHand() {
		case FiveOfAKind:
			playerByHand[0] = append(playerByHand[0], players[i])
		case FourOfAKind:
			playerByHand[1] = append(playerByHand[1], players[i])
		case FullHouse:
			playerByHand[2] = append(playerByHand[2], players[i])
		case ThreeOfAKind:
			playerByHand[3] = append(playerByHand[3], players[i])
		case TwoPair:
			playerByHand[4] = append(playerByHand[4], players[i])
		case OnePair:
			playerByHand[5] = append(playerByHand[5], players[i])
		case HighCard:
			playerByHand[6] = append(playerByHand[6], players[i])
		}
	}

	ranking := 1
	for i := len(playerByHand) - 1; i >= 0; i-- {
		playerByHand[i] = orderCards(playerByHand[i])
		totalWinnings += getWinnings(playerByHand[i], ranking)
		ranking += len(playerByHand[i])
	}

	return totalWinnings
}

// Part2 returns the answer to Day 7, Part 2 of the Advent
// of Code challenge 2023.
func Part2(filePath string) int {
	defer timer.Timer("Day 7, Part 2")()
	input, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var totalWinnings int
	var playerByHand [7][]player

	lines := strings.Split(string(input), "\n")

	players := getPlayers(lines)

	for i := 0; i < len(players); i++ {
		switch players[i].getHandWithJokers() {
		case FiveOfAKind:
			playerByHand[0] = append(playerByHand[0], players[i])
		case FourOfAKind:
			playerByHand[1] = append(playerByHand[1], players[i])
		case FullHouse:
			playerByHand[2] = append(playerByHand[2], players[i])
		case ThreeOfAKind:
			playerByHand[3] = append(playerByHand[3], players[i])
		case TwoPair:
			playerByHand[4] = append(playerByHand[4], players[i])
		case OnePair:
			playerByHand[5] = append(playerByHand[5], players[i])
		case HighCard:
			playerByHand[6] = append(playerByHand[6], players[i])
		}
	}

	ranking := 1
	for i := len(playerByHand) - 1; i >= 0; i-- {
		playerByHand[i] = orderCardsWithJokers(playerByHand[i])
		totalWinnings += getWinnings(playerByHand[i], ranking)
		ranking += len(playerByHand[i])
	}

	return totalWinnings
}

func (p *player) getHandWithJokers() hand {
	jokers := strings.Count(p.cards, "J")
	if strings.Count(p.cards, string(p.cards[0])) == 5 {
		return FiveOfAKind
	}
	for i := 0; i < len(p.cards); i++ {
		if p.cards[i] == 'J' {
			continue
		}
		if strings.Count(p.cards, string(p.cards[i])) == 4 {
			if jokers == 1 {
				return FiveOfAKind
			}
			return FourOfAKind
		}
		if strings.Count(p.cards, string(p.cards[i])) == 3 {
			if jokers == 2 {
				return FiveOfAKind
			} else if jokers == 1 {
				return FourOfAKind
			}
			for j := i + 1; j < len(p.cards); j++ {
				if strings.Count(p.cards, string(p.cards[j])) == 2 {
					return FullHouse
				}
			}
			return ThreeOfAKind
		}
		if strings.Count(p.cards, string(p.cards[i])) == 2 {
			if jokers == 3 {
				return FiveOfAKind
			} else if jokers == 2 {
				return FourOfAKind
			} else if jokers == 1 {
				for j := i + 1; j < len(p.cards); j++ {
					if strings.Count(p.cards, string(p.cards[j])) == 2 && p.cards[j] != 'J' && p.cards[i] != p.cards[j] {
						return FullHouse
					}
				}
				return ThreeOfAKind
			}
			for j := i + 1; j < len(p.cards); j++ {
				if strings.Count(p.cards, string(p.cards[j])) == 2 && p.cards[i] != p.cards[j] {
					return TwoPair
				} else if strings.Count(p.cards, string(p.cards[j])) == 3 {
					return FullHouse
				}
			}
			return OnePair
		}
	}
	if jokers == 4 {
		return FiveOfAKind
	} else if jokers == 3 {
		return FourOfAKind
	} else if jokers == 2 {
		return ThreeOfAKind
	} else if jokers == 1 {
		return OnePair
	}
	return HighCard
}

func greaterCardWithJokers(a, b rune) bool {
	order := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	aIndex := bytes.IndexRune(order, a)
	bIndex := bytes.IndexRune(order, b)
	return aIndex > bIndex
}

func orderCardsWithJokers(p []player) []player {
	sort.Slice(p, func(i, j int) bool {
		for k := 0; k < len(p[i].cards); k++ {
			if p[i].cards[k] != p[j].cards[k] {
				return greaterCardWithJokers(rune(p[i].cards[k]), rune(p[j].cards[k]))
			}
		}
		return false
	})
	return p
}
