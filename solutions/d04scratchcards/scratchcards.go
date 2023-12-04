package d04scratchcards

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

func checkCard(s string) int {
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

	total := 0
	wins := strings.Split(winning, " ")
	for _, v := range wins {
		if v == "" {
			continue
		}
		_, ok := score[v]
		if ok {
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}
	}
	return total
}

func PlayScratchcards(lines []string) int {
	total := 0
	for _, l := range lines {
		total += checkCard(l)
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(PlayScratchcards(lines)), "B"
}
