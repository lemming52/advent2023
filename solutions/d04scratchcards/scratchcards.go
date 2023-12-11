package d04scratchcards

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

type CardSet struct {
	cards []int
	total int
}

func NewCardSet(n int) *CardSet {
	c := &CardSet{
		cards: make([]int, n),
		total: 0,
	}
	for i := 0; i < n; i++ {
		c.cards[i] = 0
	}
	return c
}

func (c *CardSet) HandleCard(i, matches int) {
	c.cards[i] += 1
	for j := 1; j <= matches; j++ {
		c.cards[i+j] += c.cards[i]
	}
	c.total += c.cards[i]
}

// checkCard scores each card and returns the number of matches
func checkCard(s string) (int, int) {
	card := strings.Split(s, ":")
	numbers := strings.Split(card[1], "|")
	scoring, winning := strings.Trim(numbers[0], " "), strings.Trim(numbers[1], " ")

	scoringVals := strings.Split(scoring, " ")
	score := make(map[string]bool, len(scoringVals))
	for _, v := range scoringVals {
		if v == "" {
			continue
		}
		score[v] = true
	}

	total, count := 0, 0
	wins := strings.Split(winning, " ")
	for _, v := range wins {
		if v == "" {
			continue
		}
		_, ok := score[v]
		if ok {
			count += 1
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}
	}
	return total, count
}

func PlayScratchcards(lines []string) (int, int) {
	cards := NewCardSet(len(lines))
	total := 0
	for i, l := range lines {
		val, matches := checkCard(l)
		total += val
		cards.HandleCard(i, matches)
	}
	return total, cards.total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := PlayScratchcards(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
