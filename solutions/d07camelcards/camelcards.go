package d07camelcards

import (
	"advent/solutions/utils"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand  string
	score int
	bid   int
}

func newHand(s string) *Hand {
	components := strings.Split(s, " ")
	return &Hand{
		hand:  s,
		score: Score(components[0]),
		bid:   utils.Stoi(components[1]),
	}
}

func Score(hand string) int {
	cards := map[rune]int{}
	for _, v := range hand {
		_, ok := cards[v]
		if !ok {
			cards[v] = 1
		} else {
			cards[v] += 1
		}
	}
	highestCount, secondHighest := 0, 0
	for _, count := range cards {
		if count > highestCount {
			secondHighest = highestCount
			highestCount = count
		} else if count == highestCount {
			secondHighest = count
		} else if count > secondHighest {
			secondHighest = count
		}
	}
	return checkScore(highestCount, secondHighest)
}

func checkScore(a, b int) int {
	if a == b {
		if a == 1 {
			// High card
			return 1
		}
		// Two Pair
		return 3
	} else if a == 3 && b == 2 {
		// Full house
		return 5
	}
	switch a {
	case 4, 5:
		// 4/5 of a kind
		return a + 2
	case 3:
		// 3 of a kind
		return 4
	default:
		// pair
		return 2
	}
}

func CompareHand(a, b *Hand) bool {
	if a.score == b.score {
		return compare(a.hand, b.hand)
	}
	return a.score < b.score
}

func compare(a, b string) bool {
	for i, c := range a {
		d := rune(b[i])
		if c != d {
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				return c < d
			default:
				switch d {
				case '1', '2', '3', '4', '5', '6', '7', '8', '9', 'T':
					return false
				default:
					return compareFaceCards(c, d)
				}
			}
		}
	}
	return false
}

func compareFaceCards(a, b rune) bool {
	// We know a != b
	switch a {
	case 'T':
		return true
	case 'J':
		return b == 'Q' || b == 'K' || b == 'A'
	case 'Q':
		return b == 'K' || b == 'A'
	case 'K':
		return b == 'A'
	default:
		return false
	}
}

func ScoreHands(lines []string) int {
	hands := make([]*Hand, len(lines))
	for i, l := range lines {
		hands[i] = newHand(l)
	}
	sort.Slice(hands, func(i, j int) bool {
		return CompareHand(hands[i], hands[j])
	})
	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := ScoreHands(lines)
	return strconv.Itoa(a), "B"
}
