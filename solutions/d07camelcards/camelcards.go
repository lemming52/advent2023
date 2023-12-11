package d07camelcards

import (
	"advent/solutions/utils"
	"sort"
	"strconv"
	"strings"
)

// Hand captures the info of a hand, it's hand score/scoring rank, it's score if allowing jokers and the bid
type Hand struct {
	hand       string
	score      int
	jokerScore int
	bid        int
}

func newHand(s string) *Hand {
	components := strings.Split(s, " ")
	a, b := Score(components[0])
	return &Hand{
		hand:       s,
		score:      a,
		jokerScore: b,
		bid:        utils.Stoi(components[1]),
	}
}

// Score scores a hand, it does this by computing the counts of each card rank in the hand
func Score(hand string) (int, int) {
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
	score := checkScore(highestCount, secondHighest)
	j, ok := cards['J']
	if !ok {
		return score, score
	}
	return score, checkScoreJoker(highestCount, secondHighest, j)
}

// checkScore takes the two highest matching card counts of a hand and returns the score
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

// checkScoreJoker uses the count of jokers, and the two highest card counts to return the hand score including jokers
func checkScoreJoker(highest, second, joker int) int {
	if joker < second {
		// Two pair => Full house
		return 5
	} else if joker == second {
		if joker == 1 {
			// Always best to improve highest
			return checkScore(highest+1, second)
		}
		// Two pair or full house, one of which is jokers
		if second == highest {
			return 6
		}
		return 7
	}
	if second == 1 {
		return checkScore(highest+1, 1)
	}
	return 7

}

// CompareHand takes two hands and returns if a is worse than b
func CompareHand(a, b *Hand, joker bool) bool {
	if !joker {
		if a.score == b.score {
			return compare(a.hand, b.hand)
		}
		return a.score < b.score
	}
	if a.jokerScore == b.jokerScore {
		return compareJoker(a.hand, b.hand)
	}
	return a.jokerScore < b.jokerScore

}

// compare iterates across the hands character by character to see which is higher
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

// compareFaceCards checks which of two face cards is higher
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

// compareJoker compares two hands including joker scores
func compareJoker(a, b string) bool {
	for i, c := range a {
		d := rune(b[i])
		if c != d {
			if d == 'J' {
				return false
			}
			switch c {
			case 'J':
				return true
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

// ScoreHands ranks and totals the bids for both scores, by scoring each hand, then sorting by rank
func ScoreHands(lines []string) (int, int) {
	hands := make([]*Hand, len(lines))
	for i, l := range lines {
		hands[i] = newHand(l)
	}
	sort.Slice(hands, func(i, j int) bool {
		return CompareHand(hands[i], hands[j], false)
	})
	firstTotal := 0
	for i, h := range hands {
		firstTotal += (i + 1) * h.bid
	}

	sort.Slice(hands, func(i, j int) bool {
		return CompareHand(hands[i], hands[j], true)
	})
	secondTotal := 0
	for i, h := range hands {
		secondTotal += (i + 1) * h.bid
	}
	return firstTotal, secondTotal
}

// Run solves
// Rather than this method, changed to a list format for each hand and replaced the face cards with numerical values
// This solution is overly complicated.
func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := ScoreHands(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
